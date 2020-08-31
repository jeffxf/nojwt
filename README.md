# nojwt

Many people use JWS "JWT" to replace server side session management. They don't 
use JWE's and they'll only ever use one signing algorithm at a time. If you're 
one of these people (you probably are), you are the target audience for this 
repo.

Although this isn't endorsed by @FiloSottile (Filippo Valsorda), it's influenced 
by his twitter thread of valid JWS complaints: 
https://twitter.com/FiloSottile/status/1288964453065797632

This repo is a golang example of how you could replace JWT with Ed25519. A few 
simple wrapper functions exist in this module allowing you to create tokens 
easily.
