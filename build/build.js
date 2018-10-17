import del from 'del';
import gulp from 'gulp';
import gulpFilter from 'gulp-filter';
import gulpMinifyCss from 'gulp-minify-css';
import gulpMinifyHtml from 'gulp-minify-html';
import gulpUglify from 'gulp-uglify';
import gulpUseref from 'gulp-useref';
import gulpRev from 'gulp-rev';
import gulpRevReplace from 'gulp-rev-replace';
import gulpSize from 'gulp-size';
import uglifySaveLicense from 'uglify-save-license';
import path from 'path';

import config from "./config";

gulp.task('build', ['index:prod', 'assets'], function () {
    let htmlFilter = gulpFilter('*.html', {restore: true});
    let vendorCssFilter = gulpFilter('**/vendor.css', {restore: true});
    let vendorJsFilter = gulpFilter('**/vendor.js', {restore: true});
    let assets;

    return gulp.src(path.join(config.paths.prodTmp, '*.html'))
        .pipe(assets = gulpUseref.assets({
            searchPath: [
                // To resolve local path
                config.paths.prodTmp,
                // To resolve bower_components/... paths
                config.paths.base
            ]
        }))
        .pipe(vendorCssFilter)
        .pipe(gulpMinifyCss())
        .pipe(vendorCssFilter.restore)
        .pipe(vendorJsFilter)
        .pipe(gulpUglify({preserveComments: uglifySaveLicense}))
        .pipe(vendorJsFilter.restore)
        .pipe(gulpRev())
        .pipe(assets.restore())
        .pipe(gulpUseref({searchPath: [config.paths.prodTmp]}))
        .pipe(gulpRevReplace())
        .pipe(htmlFilter)
        .pipe(gulpMinifyHtml({
            empty: true,
            spare: true,
            quotes: true
        }))
        .pipe(htmlFilter.restore)
        .pipe(gulp.dest(config.paths.dist))
        .pipe(gulpSize({ showFiles: true }));
});

// Copies assets to the dist dir
gulp.task('assets', function () {
    return gulp.src(path.join(config.paths.assets, '/**/*'), {base: config.paths.app})
        .pipe(gulp.dest(config.paths.dist));
});

// Cleans all build artifacts
gulp.task('clean', function () {
    return del([path.join(config.paths.dist, '/'), path.join(config.paths.tmp, '/')]);
});