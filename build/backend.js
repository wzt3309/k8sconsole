import del from 'del';
import fs from 'fs';
import gulp from 'gulp';
import lodash from 'lodash';
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
            // record version info into src/app/backend/client.Version
            '-ldflags',
            config.recordVersionExpression,
            '-o',
            path.join(config.paths.serve, config.backend.binaryName),
            config.backend.mainPackageName,
        ], doneFn);
});

/**
 * Compiles backend application in production mode for the current architecture and places the
 * binary in the dist directory.
 */
gulp.task('backend:prod', ['package-backend', 'clean-dist'], function() {
    let outputBinaryPath = path.join(config.paths.dist, config.backend.binaryName);
    return backendProd([[outputBinaryPath, config.arch.default]]);
});

/**
 * Compiles backend application in production mode for all architectures and places the
 * binary in the dist directory.
 */
gulp.task('backend:prod:cross', ['package-backend', 'clean-dist'], function() {
    let outputBinaryPaths =
        config.paths.distCross.map((dir) => path.join(dir, config.backend.binaryName));
    return backendProd(lodash.zip(outputBinaryPaths, config.arch.list));
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

function backendProd(outputBinaryPathsAndArchs) {
    let promiseFn = (path, arch) => {
        return (resolve, reject) => {
            goCommand(
                [
                    'build',
                    '-a',
                    '-installsuffix',
                    'cgo',
                    '-ldflags',
                    `${config.recordVersionExpression} -w -s`,
                    '-o',
                    path,
                    config.backend.mainPackageName,
                ],
                (err) => {
                    if (err) {
                        reject(err);
                    } else {
                        resolve();
                    }
                },
                {
                    // Disable cgo package
                    CGO_ENABLED: '0',
                    GOARCH: arch,
                });
        };
    };

    let goCommandPromises = outputBinaryPathsAndArchs.map(
        (pathAndArch) => new Promise(promiseFn(pathAndArch[0], pathAndArch[1]))
    );

    return Promise.all(goCommandPromises);
}