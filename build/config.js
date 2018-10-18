import path from 'path'

/**
 * Load the i18n and l10n configuration. Used when dashboard is built in production.
 */
let localization = require('../i18n/locale_conf.json');

/**
 * project base path
 */
const basePath = path.join(__dirname, '../');

export default {

    /**
     * Absolute paths to known dirs, e.g., to src dir
     */
    paths: {
        app: path.join(basePath, 'src/app'),
        assets: path.join(basePath, 'src/app/assets'),
        backendSrc: path.join(basePath, 'src/app/backend'),
        backendTmp: path.join(basePath, '.tmp/backend'),
        backendTmpSrc: path.join(
            basePath, '.tmp/backend/src/github.com/wzt3309/k8sconsole/src/app/backend'),
        backendTmpSrcVendor: path.join(
            basePath, '.tmp/backend/src/github.com/wzt3309/k8sconsole/vendor'),
        backendVendor: path.join(basePath, 'vendor'),
        base: basePath,
        bowerComponents: path.join(basePath, 'bower_components'),
        build: path.join(basePath, 'build'),
        deploySrc: path.join(basePath, 'src/app/deploy'),
        dist: path.join(basePath, 'dist'),
        externs: path.join(basePath, 'src/app/externs'),
        frontendSrc: path.join(basePath, 'src/app/frontend'),
        frontendTest: path.join(basePath, 'src/test/frontend'),
        goTools: path.join(basePath, '.tools/go'),
        i18nProd: path.join(basePath, '.tmp/i18n'),
        integrationTest: path.join(basePath, 'src/test/integration'),
        karmaConf: path.join(basePath, 'build/karma.conf.js'),
        messagesForExtraction: path.join(basePath, '.tmp/messages_for_extraction'),
        nodeModules: path.join(basePath, 'node_modules'),
        partials: path.join(basePath, '.tmp/partials'),
        prodTmp: path.join(basePath, '.tmp/prod'),
        serve: path.join(basePath, '.tmp/serve'),
        src: path.join(basePath, 'src'),
        tmp: path.join(basePath, '.tmp'),
        xtbgenerator: path.join(basePath, '.tools/xtbgenerator/bin/XtbGenerator.jar'),
    },

    /**
     * The name of the Angular module
     */
    frontend: {
        /**
         * The name of Angular module, i.e., the module that bootstrap the application
         */
        rootModuleName: 'k8sconsole',
    },

    backend: {
        /**
         * The name of the backend binary.
         */
        binaryName: 'console',
        /**
         * Name of the main backend package that is used in go build command.
         */
        mainPackageName: 'github.com/wzt3309/k8sconsole/src/app/backend',
    },

    build: {
        prod: 'production',
        test: 'test',
        dev: 'development',
    },

    deploy: {
        imageName: 'wzt3309/k8sconsole',
    },

    /**
     * Configuration for i18n & l10n.
     */
    translations: localization.translations.map((translation) => {
        return {path: path.join(basePath, 'i18n', translation.file), key: translation.key};
    }),

};
