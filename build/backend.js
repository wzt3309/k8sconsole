import del from 'del';
import fs from 'fs';
import gulp from 'gulp';
import path from 'path';

import config from './config';
import goCommand from './gocommand';

/**
 * Compiles backend application in dev mode and places the binary in serve directory.
 */
gulp.task('backend', ['package-backend'], function (doneFn) {
    goCommand(
        [
            'build',
            // Install dependencies to speed up subsequent compilations.
            '-i',
            '-o',
            path.join(config.paths.serve, config.backend.binaryName),
            config.backend.mainPackageName,
        ], doneFn);
});

/**
 * Packages backend code to be ready for tests and compilation.
 */
gulp.task('package-backend', ['package-backend-source', 'link-vendor']);

/**
 * Moves all backend source files (app and tests) to a temporary package directory where it can be
 * applied go commands.
 */
gulp.task('package-backend-source', ['clean-packaged-backend-source'], function () {
    return gulp.src([path.join(config.paths.backendSrc, '**/*')])
        .pipe(gulp.dest(config.paths.backendTmpSrc));
});


/**
 * Cleans packaged backend source to remove any leftovers from there.
 */
gulp.task('clean-packaged-backend-source', function() {
    return del([config.paths.backendTmpSrc]);
});

/**
 * Links vendor folder to the packaged backend source
 */
gulp.task('link-vendor', ['package-backend-source'], function (doneFn) {
    fs.symlink(config.paths.backendVendor, config.paths.backendTmpSrcVendor, 'dir', (err) => {
        if (err && err.code === 'EEXIST') {
            // Skip errors if the link already exists.
            doneFn();
        } else {
            doneFn(err);
        }
    });
});