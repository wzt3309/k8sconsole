require('babel-core/register');
var conf = require('./config');
var path = require('path');


exports.config = {
    capabilities: {
        'browserName': 'chrome',
    },

    baseUrl: 'http://localhost:3000',

    specs: [path.join(conf.paths.integrationTest, '**/*.js')],
};
