import browserSync from 'browser-sync';
import gulp from 'gulp';
import gulpProtractor from 'gulp-protractor';
import karma from 'karma';
import path from 'path';

import config from './config';

/**
 *
 * @param {boolean} singleRun
 * @param {function(?Error=)} doneFn
 */
function runUnitTests(singleRun, doneFn) {
    let localConfig = {
        configFile: config.paths.karmaConf,
        singleRun: singleRun,
        autoWatch: !singleRun,
    };

    let server = new karma.Server(localConfig, function (failCount) {
        doneFn(failCount ? new Error("Failed" + failCount + " tests.") : undefined);
    });
    server.start();
}

function runProtractorTests(doneFn) {
    gulp.src(path.join(config.paths.integrationTest, '**/*.js'))
        .pipe(gulpProtractor.protractor({
            configFile: config.paths.protractorConf,
        }))
        .on('error', function (err) {
            doneFn(err);
        })
        .on('end', function () {
            // Close browser sync server.
            browserSync.exit();
            doneFn();
        });
}

// Run once all unit tests of the application
gulp.task('test', function (doneFn) {
    runUnitTests(true, doneFn);
});

// Runs all unit tests of the application. Watches for changes in the source files to rerun
// the tests.

gulp.task('test:watch', function (doneFn) {
    runUnitTests(false, doneFn);
});

/**
 * Runs application integration tests. Uses development version of the application.
 */
gulp.task('integration-test', ['serve', 'webdriver-update'], runProtractorTests);


/**
 * Runs application integration tests. Uses production version of the application.
 */
gulp.task('integration-test:prod', ['serve:prod', 'webdriver-update'], runProtractorTests);


/**
 * Downloads and updates webdriver. Required to keep it up to date.
 */
gulp.task('webdriver-update', gulpProtractor.webdriver_update);

