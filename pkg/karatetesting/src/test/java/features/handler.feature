Feature: work with DB again
  
Background: connect to Db again
    * def dbHandler = Java.type('helpers.DbHandler')
    * configure afterScenario = 
      """
      function(){
      karate.log('after scenario:', karate.scenario.name);
      }
      """  
    * configure afterFeature = 
      """
      function(){
      karate.log('after feature:', karate.feature.name);
      }
      """  
@smoke
Scenario: Create a pizza
    * def pizzaName = "PizzaTestAGABA"
    * def price = 320
    * def sqlAddPizza = "INSERT INTO pizzas (name, price) VALUES ('" + pizzaName + "', " + price + ")"
    * print sqlAddPizza
    # * eval dbHandler.postData(sqlAddPizza)

Scenario: Create a Topping
  * def toppingName = "toppingtestGABA"
  * def price = 400
  * def is_internal = false
  * def sqlAddTopping = "INSERT INTO toppings (name, price, is_internal) VALUES ('" + toppingName + "', " + price + ", " + is_internal + ")"
  * print sqlAddTopping
  # * eval dbHandler.postData(sqlAddTopping)
@smoke
Scenario: Get a Pizza by Ids
    * def level = dbHandler.getPizzasById(1)
    * print level.price
    * print level.name
    * print "Executing and Testing"
    * print "name---",myVarName

# Scenario: Call another feature file
#     * call read('my.feature'){message: "Heyyy"}