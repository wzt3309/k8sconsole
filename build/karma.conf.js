import path from 'path'
import wiredep from 'wiredep';

import config from './config'

/**
 * Returns an array of files required by Karma to run the tests.
 * @returns {!Array<string>}
 */
function getFileList() {
    // All app dependencies are required for tests. Include them.
    let wiredepOptions = {
        dependencies: true,
        devDependencies: true
    };

    return wiredep(wiredepOptions).js
        .concat([
            path.join(config.paths.frontendTest, '**/*.js'),
            path.join(config.paths.frontendSrc, '**/*.js'),
            path.join(config.paths.frontendSrc, '**/*.html')
        ]);
}

/**
 * Exported default function which sets Karma configuration. Required by the framework.
 *
 * @param {!Object} conf
 */
export default function (conf) {
    let configuration = {
        basePath: config.paths.base,

        files: getFileList(),

        frameworks: ['jasmine', 'browserify'],

        browsers: ['Chrome'],

        reporters: ['progress'],

        preprocessors: {}, // This field is filled with values later.

        plugins: [
            'karma-chrome-launcher',
            'karma-jasmine',
            'karma-ng-html2js-preprocessor',
            'karma-sourcemap-loader',
            'karma-browserify'
        ],

        // karma-browserify plugin config.
        browserify: {
            // Add source maps to outpus bundles.
            debug: true,
            // Make 'import ...' statements relative to the following paths.
            paths: [config.paths.frontendSrc, config.paths.frontendTest],
            transform: [
                // Transform ES6 code into ES5 so that browsers can digest it.
                'babelify'
            ]
        },

        // karma-ng-html2js-preprocessor plugin config.
        ngHtml2JsPreprocessor: {
            stripPrefix: config.paths.frontendSrc + '/',
            moduleName: config.frontend.moduleName
        }
    };

    // Convert all JS code written ES6 with modules to ES5 bundles that browsers can digest.
    configuration.preprocessors[path.join(config.paths.frontendTest, '**/*.js')] = ['browserify'];
    configuration.preprocessors[path.join(config.paths.frontendSrc, '**/*.js')] = ['browserify'];

    // Convert HTML templates into JS files that serve code through $templateCache.
    configuration.preprocessors[path.join(config.paths.frontendSrc, '**/*.html')] = ['ng-html2js'];

    conf.set(configuration);
}