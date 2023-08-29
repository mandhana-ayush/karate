Feature: work with DB again
  
Background: connect to Db again
    * def dbHandler = Java.type('helpers.DbHandler')

Scenario: Create a pizza
  # Given def receivedMessage = message
  * print 'opd Message:'