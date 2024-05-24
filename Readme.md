# Currency translator
Simple app translating currency provided in float to text.

## Redis test
Initial creation of text string is about 100us (medium 60us).
When using Redis for cache this timing is:
* with saving ~2.8ms
* with reading ~2ms

Summary: It's best to use Redis cache when calculation/operations are over 3ms 

## Initial work
Greg, Lord of Mailgun Messages, Master of Redis Realms, Messenger of Symfony Secrets, Champion of Domain-Driven Development, Guardian of PHPUnit Proclamations, and Sage of PHP Sorcery