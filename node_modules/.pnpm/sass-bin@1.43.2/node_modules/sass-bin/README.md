## Sass Dart-VM-based binary node module

This package bundles the Dart-VM based binary version of Dart Sass for use through the `npm` package manager.

Versionining follows the major and minor version of the original `sass` package and the latest available patch version.

Therefore you should include `sass-bin` like `^1.42.1` or `~1.42.1` and know that the patch version will not match the patch version of `sass`.

You can see performance increases of around **+500%** in typical Sass projects. This is because the Dart VM executes the Dart implementation much faster than the JS compiled version of Dart Sass.

Read more here about the details of the performance difference: [https://sass-lang.com/dart-sass](https://sass-lang.com/dart-sass)

### How to install

`npm` link: [https://www.npmjs.com/package/sass-bin](https://www.npmjs.com/package/sass-bin)

Use `sass-bin` instead of `sass` in your node module `package.json` dependencies. 

### Versioning

The versioning of this package follows the version of the official implementation: [https://github.com/sass/dart-sass](https://github.com/sass/dart-sass)

### macOS / M1

`arch -x86_64 zsh -c "npm install sass-bin"` can be used to run the x64 darwin binary on the new M1 processor. 
