/**
 * @fileoverview Gulp tasks that index files with dependencies (e.g., CSS or JS) injected.
 */
import browserSync from 'browser-sync';
import fs from 'fs';
import gulp from 'gulp';
import gulpInject from 'gulp-inject';
import path from 'path';
import wiredep from 'wiredep';

import config from './config';

/**
 * Creates index file in the given directory with dependencies injected from that directory.
 *
 * @param {string} indexPath
 * @param {boolean} dev
 * @return {!stream.Stream}
 */
function createIndexFile(indexPath, dev) {
    let injectStyles = gulp.src(path.join(indexPath, '**/*.css'), {read: false});
    let injectScripts = gulp.src(path.join(indexPath, '**/*.js'), {read: false});
    let injectOptions = {
        // Make the dependencies relative to the deps directory.
        ignorePath: [path.relative(config.paths.base, indexPath)],
        addRootSlash: false,
        quiet: true,
    };

    let wiredepOptions = {
        // Make wiredep dependencies begin with "node_modules/" not "../../...".
        ignorePath: path.relative(config.paths.frontendSrc, config.paths.base),
        bowerJson: JSON.parse(fs.readFileSync(path.join(config.paths.base, 'package.json'))),
        directory: config.paths.nodeModules,
        devDependencies: false,
        customDependencies: ['easyfont-roboto-mono'],
        onError: (msg) => {
            console.log(msg);
        },
    };

    if (dev) {
        wiredepOptions.customDependencies.push('google-closure-library');
    }


    return gulp.src(path.join(config.paths.frontendSrc, 'index.html'))
        .pipe(gulpInject(injectStyles, injectOptions))
        .pipe(gulpInject(injectScripts, injectOptions))
        .pipe(wiredep.stream(wiredepOptions))
        .pipe(gulp.dest(indexPath));
}

/**
 * Creates frontend application index file with development dependencies injected.
 */
gulp.task('index', ['scripts', 'styles'], function () {
    return createIndexFile(config.paths.serve, true).pipe(browserSync.stream());
});

/**
 * Creates frontend application index file with production dependencies injected.
 */
gulp.task('index:prod', ['scripts:prod', 'styles:prod'], function () {
    return createIndexFile(config.paths.prodTmp, false);
});
