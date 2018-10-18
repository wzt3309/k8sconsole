import gulp from 'gulp';
import gulpAutoprefixer from 'gulp-autoprefixer';
import gulpConcat from 'gulp-concat';
import gulpMinifyCss from 'gulp-minify-css';
import gulpSass from 'gulp-sass';
import gulpSourcemaps from 'gulp-sourcemaps';

import path from 'path';

import config from './config';

/**
 * Compile stylesheets and places them into serve folder. Each stylesheet file is compiled separately
 */
gulp.task('styles', function () {
    let sassOptions = {
        style: 'expanded',
    };

    return gulp.src(path.join(config.paths.frontendSrc, '**/*.scss'))
        .pipe(gulpSourcemaps.init())        // mark the init point of source map
        .pipe(gulpSass(sassOptions))
        .pipe(gulpAutoprefixer())
        .pipe(gulpSourcemaps.write('.'))    // end of source map
        .pipe(gulp.dest(config.paths.serve))
});

/**
 * Compile stylesheets and places them into the prod tmp folder. Stylesheets are compiled and minified
 * into a single file
 */
gulp.task('style:prod', function () {
    let sassOptions = {
        style: 'compressed',
    };

    return gulp.src(path.join(config.paths.frontendSrc, '**/*.scss'))
        .pipe(gulpSass(sassOptions))
        .pipe(gulpAutoprefixer())
        .pipe(gulpConcat('app.css'))
        .pipe(gulpMinifyCss({
            // Do not process @import statements. This breaks Angular Material font icons.
            processImport: false,
        }))
        .pipe(gulp.dest(config.paths.prodTmp));
});
