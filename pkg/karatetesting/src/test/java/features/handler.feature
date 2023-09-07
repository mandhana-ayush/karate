Feature: work with DB again
  
Background: connect to Db again
    * def dbHandler = Java.type('helpers.DbHandler')
    # * def tokenFetch = Java.type('helpers.TokenFetch')
    # * def v = tokenFetch.getAccessToken(url, refreshToken)
    # * print v.data.access_token
    # * configure afterScenario = 
    #   """
    #   function(){
    #   karate.log('after scenario:', karate.scenario.name);
    #   }
    #   """  
    # * configure afterFeature = 
    #   """
    #   function(){
    #   karate.log('after feature:', karate.feature.name);
    #   }
    #   """  
    * def sqlAddPizza = ""
@smoke
Scenario: Create a pizza
    * def pizzaName = "PizzaTestAGABA"
    * def price = 320
    * eval sqlAddPizza = "INSERT INTO pizzas (name, price) VALUES ('" + pizzaName + "', " + price + ")"
    * print accessToken
    # * print sqlAddPizza
    # * eval dbHandler.postData(sqlAddPizza)

Scenario: Create a Topping
  # * print sqlAddPizza
  * def toppingName = "toppingtestGABA"
  * def price = 400
  * def is_internal = false
  * def sqlAddTopping = "INSERT INTO toppings (name, price, is_internal) VALUES ('" + toppingName + "', " + price + ", " + is_internal + ")"
  * print accessToken
  # * print sqlAddTopping
  # * eval dbHandler.postData(sqlAddTopping)
@smoke
Scenario: Get a Pizza by Ids
    * def level = dbHandler.getPizzasById(1)
    * print accessToken
    # * print level.price
    # * print level.name
    # * print "Executing and Testing"
    # * print "name---",myVarName

# Scenario: calling two api request
#     # First API request with path '/endpoint1'
#   Given path '/ping'
#   When method GET
#   Then status 200
#   And response.message == "pong"

#   # Second API request with path '/endpoint2'
#   Given path '/piang'
#   When method GET
#   Then status 200
#   And response.message == "poang"
# # Scenario: Call another feature file
# #     * call read('my.feature'){message: "Heyyy"}