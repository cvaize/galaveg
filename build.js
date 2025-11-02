const fs = require('fs');
const zlib = require('zlib');
const minify = require('@node-minify/core');
const cleanCSS = require('@node-minify/clean-css');
const uglifyjs = require('@node-minify/uglify-js');

const styles = [
    ...fs.readdirSync('./resources/css/libraries')
        .filter(s => s.endsWith('.css'))
        .map(s => './resources/css/libraries/' + s),

    './resources/css/components/layout.css',

    ...fs.readdirSync('./resources/css/components')
        .filter(s => s.endsWith('.css') && s !== 'layout.css')
        .map(s => './resources/css/components/' + s),
].filter((item, i, ar) => ar.indexOf(item) === i);

const scripts = [
    ...fs.readdirSync('./resources/js/libraries')
        .filter(s => s.endsWith('.js'))
        .map(s => './resources/js/libraries/' + s),

    ...fs.readdirSync('./resources/js/components')
        .filter(s => s.endsWith('.js'))
        .map(s => './resources/js/components/' + s),
].filter((item, i, ar) => ar.indexOf(item) === i);

const svg = [
    ...fs.readdirSync('./resources/svg')
        .filter(s => s.endsWith('.svg'))
        .map(s => './resources/svg/' + s),
].filter((item, i, ar) => ar.indexOf(item) === i);

async function runStyles(){
    !fs.existsSync("static/css") && fs.mkdirSync("static/css")
    let content = '';

    for (const style of styles) {
        content += fs.readFileSync(style);
    }

    fs.writeFileSync('./static/css/app.css', content);

    await minify({
        compressor: cleanCSS,
        input: './static/css/app.css',
        output: './static/css/app.min.css'
    });

    content = fs.readFileSync('./static/css/app.min.css');

    content = zlib.gzipSync(content, {level: 9});

    fs.writeFileSync('./static/css/app.min.css.gz', content);

}

async function runScripts(){
    !fs.existsSync("static/js") && fs.mkdirSync("static/js")
    let content = '';

    for (const script of scripts) {
        content += fs.readFileSync(script);
    }

    fs.writeFileSync('./static/js/app.js', content);

    await minify({
        compressor: uglifyjs,
        input: './static/js/app.js',
        output: './static/js/app.min.js',
    });

    content = fs.readFileSync('./static/js/app.min.js');

    content = zlib.gzipSync(content, {level: 9});

    fs.writeFileSync('./static/js/app.min.js.gz', content);
}
async function runSvg(){
    !fs.existsSync("static/svg") && fs.mkdirSync("static/svg")
    for (const svgElement of svg) {
        let content = fs.readFileSync(svgElement);
        let path = svgElement.replace('./resources/svg/', './static/svg/').trim();
        fs.writeFileSync(path, content);

        content = zlib.gzipSync(content, {level: 9});

        fs.writeFileSync(`${path}.gz`, content);
    }

}

async function run(){
    await runStyles();
    await runScripts();
    await runSvg();
}

run();
