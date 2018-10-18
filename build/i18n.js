/**
 * Gulp tasks for processing and compiling frontend JavaScript files.
 */
import childProcess from 'child_process';
import fileExists from 'file-exists';
import gulp from 'gulp';
import cheerio from 'gulp-cheerio';
import freplace from 'gulp-findreplace';
import gulpUtil from 'gulp-util';
import jsesc from 'jsesc';
import path from 'path';
import q from 'q';
import regexpClone from 'regexp-clone';

import config from './config';

function extractForLanguage(langKey) {
    let deferred = q.defer();

    let translationBundle = path.join(config.paths.base, `i18n/messages-${langKey}.xtb`);
    let codeSource = path.join(config.paths.serve, '**.js');
    let messagesSource = path.join(config.paths.messagesForExtraction, '**.js');
    let command = `java -jar ${config.paths.xtbgenerator} --lang ${langKey}` +
        ` --xtb_output_file ${translationBundle} --js ${codeSource} --js ${messagesSource}`;
    if (fileExists.sync(translationBundle)) {
        command = `${command} --translations_file ${translationBundle}`;
    }

    childProcess.exec(command, function(err, stdout, stderr) {
        if (err) {
            gulpUtil.log(stdout);
            gulpUtil.log(stderr);
            deferred.reject(new Error(err));
        }
        deferred.resolve();
    });

    return deferred.promise;
}


gulp.task('generate-xtbs', [
    'extract-translations',
    'remove-unused-translations',
    'set-prod-node-env',
]);

let prevMsgs = {};

gulp.task('buildExistingI18nCache', function() {
    return gulp.src('i18n/messages-en.xtb').pipe(cheerio((doc) => {
        doc('translation').each((i, translation) => {
            let key = translation.attribs.key;
            let index = key.lastIndexOf('_');
            if (index !== -1) {
                let lastpart = key.substring(index + 1);
                if (/^[0-9]+$/.test(lastpart)) {
                    let indexSuffix = Number(lastpart);
                    let filePrefix = key.substring(0, index);
                    if (!prevMsgs[filePrefix]) {
                        prevMsgs[filePrefix] = [];
                    }
                    prevMsgs[filePrefix].push({
                        index: indexSuffix,
                        key: key,
                        text: doc(translation).text(),
                        desc: translation.attribs.desc,
                        used: false,
                    });
                }
            }
        });
    }));
});

/**
 * Extracts all translation messages into XTB bundles.
 *
 * Cleans up the data from the previous run to prevent cross-pollination between branches.
 */
gulp.task(
    'extract-translations', ['scripts', 'angular-templates', 'clean-messages-for-extraction'],
    function() {
        let promises = config.translations.map((translation) => extractForLanguage(translation.key));
        return q.all(promises);
    });

/**
 * Task to used to find translations used in JavaScript files. Do not run manually. It should be
 * invoked as a part of 'remove-unused-translations'.
 */
gulp.task('find-translations-used-in-js', function() {
    let jsSource = path.join(config.paths.frontendSrc, '**/*.js');
    return gulp.src(jsSource).pipe(freplace(/MSG_\w*/g, function(match) {
        // Mark every message found in JavaScript files as used, it will allow deletion of unused
        // messages afterwards.
        translationsManager.addUsed(match);
    }));
});

/**
 * Task to remove unused translations. Do not run manually. It should be invoked as a part of
 * 'generate-xtbs'.
 */
gulp.task(
    'remove-unused-translations', ['angular-templates', 'find-translations-used-in-js'],
    function() {
        // Get translations used in JavaScript and HTML files. These will not be removed.
        let used = translationsManager.getUsed();

        return gulp.src('i18n/messages-*.xtb')
            .pipe(cheerio((doc) => {
                let unused = new Set();

                // Find translations to remove.
                doc('translation').each((i, translation) => {
                    let key = translation.attribs.key;
                    if (!used.has(key)) {
                        unused.add(key);
                    }
                });

                // Remove unused translations.
                unused.forEach((r) => {
                    doc(`translation[key=${r}]`).remove();
                });
            }))
            .pipe(gulp.dest('i18n/'));
    });

/**
 * Translations manager is a closure function allowing to manage translations.
 * Allows removing unused translations after marking certain as used.
 *
 * @type {{markAsUsed, removeUnused}}
 */
export let translationsManager = (function() {

    /**
     * Set of translations marked as used.
     *
     * @type {Set}
     */
    let used = new Set();

    /**
     * Function used to mark translations as used.
     *
     * @param {string} key
     */
    function addUsed(key) {
        used.add(key);
    }

    /**
     * Function used to get used translations.
     *
     * @return {Set}
     */
    function getUsed() {
        return used;
    }

    return {
        addUsed: addUsed,
        getUsed: getUsed,
    };
})();

// Regex to match [[Foo | Bar]] or [[Foo]] i18n placeholders.
// Technical details:
// * First capturing group is lazy math for any string not-containing |. This is to make
//   both [[ message | description ]] and [[ message ]] work.
// * Second is non-capturing and optional. It has a capturing group inside. This is to
//   extract description that is optional.
const I18N_REGEX = /\[\[([^|]*?)(?:\|(.*?))?\]\]/g;

export function processI18nMessages(file, minifiedHtml) {
    let content = jsesc(minifiedHtml);
    let pureHtmlContent = `${content}`;
    let filePath = path.relative(file.base, file.path);
    let messageVarPrefix = filePath.toUpperCase().split('/').join('_').replace('.HTML', '');
    let used = new Set();

    /**
     * Finds all i18n messages inside a template and returns its text, description and original
     * string.
     * @param {string} htmlContent
     * @return {!Array<{text: string, desc: string, original: string}>}
     */
    function findI18nMessages(htmlContent) {
        let matches = htmlContent.match(I18N_REGEX);
        if (matches) {
            return matches.map((match) => {
                let exec = regexpClone(I18N_REGEX).exec(match);
                // Default to no description when it is not provided.
                let desc = (exec[2] || '(no description provided)').trim();
                // replace {{$variableName}} with {{ $variableName}} to avoid {$ getting recognised as
                // google.getMsg format
                let text = exec[1].replace('{$', '{ $');
                let varName = undefined;
                if (prevMsgs[`MSG_${messageVarPrefix}`]) {
                    for (let msg of prevMsgs[`MSG_${messageVarPrefix}`]) {
                        if (msg.text === text && msg.desc === desc && !msg.used) {
                            varName = msg.key;
                            msg.used = true;
                            used.add(msg.index);
                            break;
                        }
                    }
                }
                return {text: text, desc: desc, original: match, varName: varName};
            });
        }
        return [];
    }

    let i18nMessages = findI18nMessages(content);

    /**
     * @param {number} index
     * @return {string}
     */
    function createMessageVarName() {
        for (let i = 0;; i++) {
            if (!used.has(i)) {
                used.add(i);
                return `MSG_${messageVarPrefix}_${i}`;
            }
        }
    }

    i18nMessages.forEach((message) => {
        message.varName = message.varName || createMessageVarName();
        // Replace i18n messages with english messages for testing and MSG_ vars invocations
        // for compiler passses.
        content = content.replace(message.original, `' + ${message.varName} + '`);
        pureHtmlContent = pureHtmlContent.replace(message.original, message.text);

        // Mark every message found in this HTML file as used, it will allow deletion of unused messages
        // afterwards.
        translationsManager.addUsed(message.varName);
    });

    let messageVariables = i18nMessages.map((message) => {
        return `/** @desc ${message.desc} */\n` +
            `var ${message.varName} = goog.getMsg('${message.text}');\n`;
    });

    file.messages = messageVariables.join('\n');
    // Eval pure HTML content, because it has been jsescaped previously. This is safe to eval since
    // it was escaped by jsecs previously.
    file.pureHtmlContent = eval(`'${pureHtmlContent}'`);
    file.moduleContent = `` +
        `import module from '/index_module';\n\n${file.messages}\n` +
        `module.run(['$templateCache', ($templateCache) => {\n` +
        `    $templateCache.put('${filePath}', '${content}');\n` +
        `}]);\n`;

    return minifiedHtml;
}