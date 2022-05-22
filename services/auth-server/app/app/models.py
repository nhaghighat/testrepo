from sqlalchemy.sql import text

from app.connections import db


class User(db.Model):

    __tablename__ = "users"

    uuid = db.Column(
        db.String(36),
        primary_key=True,
        # required to be referenced as a foreign key by the Token table,
        # since postgres does not accept foreign keys referencing non
        # unique fields
        unique=True,
    )

    username = db.Column(
        db.String(255),
        primary_key=True,
    )

    password_hash = db.Column(db.String(255), nullable=False)

    token_hash = db.Column(db.String(255))

    is_admin = db.Column(db.Boolean, server_default=text("False"), nullable=False)

    created = db.Column(
        db.DateTime,
        unique=False,
        nullable=False,
        server_default=text("timezone('utc', now())"),
    )


class Token(db.Model):

    __tablename__ = "tokens"

    token = db.Column(db.String(255))

    user = db.Column(
        db.String(36), db.ForeignKey("users.uuid", ondelete="CASCADE"), primary_key=True
    )

    created = db.Column(
        db.DateTime,
        unique=False,
        nullable=False,
        server_default=text("timezone('utc', now())"),
    )
