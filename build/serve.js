import browserSync from 'browser-sync';
import browserSyncSpa from 'browser-sync-spa';
import gulp from 'gulp';
import path from 'path';

import config from "./config";

function browserSyncInit(baseDir) {
    browserSync.use(browserSyncSpa({
        selector: '[ng-app]',
    }));

    browserSync.instance = browserSync.init({
        startPath: '/',
        server: {
            baseDir: baseDir,
        },
        browser: [],
    });
}

// Serve application in dev mode.
gulp.task('serve', ['watch'], function () {
    browserSyncInit([
        config.paths.serve,
        config.paths.frontendSrc,
        config.paths.app,
        config.paths.base,
    ]);
});

// Serve application in prod mode.
gulp.task('serve:prod', ['build'], function () {
    browserSyncInit(config.paths.dist);
});

// Watch for changes in source files and run gulp tasks to rebuild them.
gulp.task('watch', ['index'], function () {
    gulp.watch([path.join(config.paths.frontendSrc, 'index.html'), 'bower.json'], ['index']);

    gulp.watch([
        path.join(config.paths.frontendSrc, '**/*.scss'),
    ], function (event) {
        if (event.type === 'changed') {
            // If is file changed, rebuild style files
            gulp.start('styles');
        } else {
            // If is file new/del, everything has to rebuilt
            gulp.start('index');
        }
    });

    gulp.watch([path.join(config.paths.frontendSrc, '**/*/js')], ['scripts'])
});
