Troubleshooting
===============

When running against issues it can be helpful to increase the verbosity of Orchest and
changing the log level of all Orchest's containers. You can do this using:

.. code-block:: sh

   orchest patch --log-level=DEBUG


Some other kubectl commands that can be useful when debugging Orchest:

.. code-block:: sh

   # Inspect the logs of a particular service
   kubectl logs -n orchest -f deployment/orchest-api

   # Attach a shell in a particular service
   kubectl exec -n orchest -it deployment/orchest-api bash

Exit code ``137`` when building Orchest images (scripts/build_containers.sh)
----------------------------------------------------------------------------
For Docker Desktop users, make sure increase the allocated memory to Docker Engine. This can be done
by going to Docker Desktop > Settings > Advanced > Increase the *Memory* slider (`GitHub issue for
reference <https://github.com/moby/moby/issues/22211>`_).

Inspecting the ``orchest-api`` API schema
-----------------------------------------
To develop against the API it can be useful to have a look at the swagger documentation. This can be
done by portforwarding the ``orchest-api`` and visiting the `/api` endpoint.

.. code-block:: sh

   # You will be able to visit `localhost:8000/api`
   kubectl port-forward -n orchest deployment/orchest-api 8000:80

Inspecting the ``orchest-database``
----------------------------------------------
.. code-block:: sh

   kubectl port-forward -n orchest deployment/orchest-database 5432:5432

   # You could accomplish the same by ``exec``ing into the database pod,
   # this can be much more handy since commands history will be
   # preserved through restarts, etc.
   psql -h 127.0.0.1 -p 5432 -U postgres -d orchest_api
   psql -h 127.0.0.1 -p 5432 -U postgres -d orchest_webserver
   psql -h 127.0.0.1 -p 5432 -U postgres -d auth_server

Breaking schema changes
-----------------------
**What it looks like**
    The client can't be accessed (the webserver is not up) or the client can be accessed but a lot
    of functionality seems to not be working, e.g. creating an environment.

**How to solve**
    .. code-block:: bash

      kubectl port-forward -n orchest deployment/orchest-database 5432:5432
      psql -h 127.0.0.1 -p 5432 -U postgres
      # Once in psql, drop the db of interest.
      drop database orchest_api; # or orchest-webserver, auth-server
      # Exit psql and restart Orchest
      bash orchest restart

    .. note::

      An alternative approach is to reinstall Orchest. ``bash orchest uninstall``
      followed by `bash orchest install``.

**Context**
    Some branches might contain a schema migration that applies changes to the database in a way
    that is not compatible with ``dev`` or any other branch. By moving back to those branches, the
    database has a schema that is not compatible with what's in the code.

**Verify**
    Check the webserver and the api logs. It will be easy to spot because the service won't produce other logs
    but the ones related to incompatible schema changes and database issues.

Error: Multiple head revisions
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
**What it looks like**
    You see an error along the lines of ``Error: Multiple head revisions are present for given
    argument 'head'`` inside one of the services interacting with the DB, e.g. the ``orchest-api``.

**How to solve**
    Using the ``orchest-api`` as an example here.

    .. code-block:: bash

       bash scripts/migration_manager.sh orchest-api merge heads

    It may be that the above doesn't work, because the ``orchest-api`` never reaches a running
    state. In that case we need to:

    .. code-block:: bash

       # Change the deployment so that it does a sleep instead of invoke
       # the cmd of the container.
       kubectl -n orchest edit deploy orchest-api
       # command: ["sleep"]
       # args: ["1000"]

       # Now run the migration script inside the orchest-api container
       python migration_manager.py db merge heads

       # Next we need to copy the file out of the container and into
       # the migration revisions directly inside the orchest-api
       kubectl cp \
           "orchest/${pod_name}:/orchest/services/orchest-api/app/migrations/versions" \
           "services/orchest-api/app/migrations/versions"

       # Rebuild the orchest-api container on the node
       scripts/build_container.sh -i orchest-api -t "v2022.04.0" -o "v2022.04.0"

       # Edit the orchest-api deployment again to make sure to not
       # run the sleep command anymore.
       kubectl -n orchest edit deploy orchest-api


**Context**
    Alembic creates revision files to do migrations. When two different branches have done schema
    migrations then the head will diverge, similar to git now having two different branches which
    point to different commits. Once these branches get merged, the alembic revision heads need to
    be merged as well.

Dev mode not working
--------------------
* Make sure you started the cluster with the Orchest repository mounted, see :ref:`here <cluster mount>`.
* If you have changed some dependencies (i.e. requirements.txt files) you have to rebuild the image and
  kill the pod to get it redeployed.

