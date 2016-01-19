'use strict';
var async = require('asyncawait/async');
var await = require('asyncawait/await');
var Promise = require('bluebird');

var request = Promise.promisifyAll(require('request'));

var getQuote = async(function () {

    let quote = await(request.getAsync('http://ron-swanson-quotes.herokuapp.com/quotes'));

    return quote;
})

function main() {
    getQuote().then(function(quote){
        console.log(quote);
    });
}

main();
