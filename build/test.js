import config from './config';
import gulp from 'gulp';
import karma from 'karma';

/**
 *
 * @param {boolean} singleRun
 * @param {function(?Error=)} doneFn
 */
function runUnitTests(singleRun, doneFn) {
    let localConfig = {
        configFile: config.paths.karmaConf,
        singleRun: singleRun,
        autoWatch: !singleRun
    };

    let server = new karma.Server(localConfig, function (failCount) {
        doneFn(failCount ? new Error("Failed" + failCount + " tests.") : undefined);
    });
    server.start();
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
