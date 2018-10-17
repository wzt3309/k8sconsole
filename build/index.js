import browserSync from 'browser-sync';
import gulp from 'gulp';
import gulpInject from 'gulp-inject';
import path from 'path';
import wiredep from 'wiredep';

import config from 'config';

function createIndexFile(indexPath) {
    let injectStyles = gulp.src(path.join(indexPath, '**/*.css'), {read: false});

    let injectScripts = gulp.src(path.join(indexPath, '**/*.js'), {read: false});

    let injectOptions = {
        ignorePath: [path.relative(config.paths.base, indexPath)],
        addRootSlash: false,
        quiet: true
    };

    let wiredepOptions = {
        ignorePath: path.relative(config.paths.frontendSrc, indexPath) + '/'
    };

    return gulp.src(path.join(config.paths.frontendSrc, 'index.html'))
        .pipe(gulpInject(injectStyles, injectOptions))
        .pipe(gulpInject(injectScripts, injectOptions))
        .pipe(wiredep.stream(wiredepOptions))
        .pipe(gulp.dest(indexPath))
        .pipe(browserSync.stream());
}

// Create frontend index file with dev deps injected
gulp.task('index', ['scripts', 'styles'], function () {
    return createIndexFile(config.paths.serve);
});

// Create frontend index file with prod deps injected
gulp.task('index:prod', ['scripts:prod', 'styles:prod'], function () {
    return createIndexFile(config.paths.prodTmp);
});