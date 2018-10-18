/**
 * @fileoverview Gulp tasks for building the project.
 */
import del from 'del';
import gulp from 'gulp';
import path from 'path';

import config from './config';

// Cleans all build artifacts
gulp.task('clean', function () {
    return del([path.join(config.paths.dist, '/'), path.join(config.paths.tmp, '/')]);
});