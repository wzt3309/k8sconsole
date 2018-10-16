import path from 'path'

/**
 * project base path
 */
const basePath = path.join(__dirname, '../');

export default {

    /**
     * Absolute paths to known dirs, e.g., to src dir
     */
    paths: {
        app: path.join(basePath, 'src/app'),
        assets: path.join(basePath, 'src/app/assets'),
        backendSrc: path.join(basePath, 'src/app/backend'),
        base: basePath,
        bowerComponents: path.join(basePath, 'bower_components'),
        build: path.join(basePath, 'build'),
        dist: path.join(basePath, 'dist'),
        frontendSrc: path.join(basePath, 'src/app/frontend'),
        frontendTest: path.join(basePath, 'src/test/frontend'),
        src: path.join(basePath, 'src'),
        tmp: path.join(basePath, '.tmp')
    }
};
