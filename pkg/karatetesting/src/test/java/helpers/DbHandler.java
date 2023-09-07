package helpers;

import java.sql.Connection;
import java.sql.DriverManager;
import java.sql.ResultSet;
import java.sql.SQLException;
import java.sql.Statement;
import java.util.List;

import com.oracle.truffle.regex.tregex.util.json.JsonObject;

import net.minidev.json.JSONObject;

public class DbHandler {
  public static String connectionUrl = "jdbc:postgresql://localhost:5433/pizza_shop_test?user=postgres&password=Manushya@123";

  public static String connectionCoreUrl = "jdbc:postgresql://localhost:5433/postgres?user=postgres&password=Manushya@123";

  public static JSONObject getPizzasById(int id) {
    JSONObject json = new JSONObject();

    try (Connection connect = DriverManager.getConnection(connectionUrl)) {
      ResultSet rs = connect.createStatement().executeQuery("SELECT * FROM pizzas WHERE id ='" + id + "'");
      // rs.next();

      if (rs.next()) {
        json.put("name", rs.getString("name"));
        json.put("price", rs.getString("price"));
      } else {
        // Handle case where no row is found for the given id
        json.put("error", "Pizza not found");
      }
      connect.close();
      return json;
    } catch (SQLException e) {
      e.printStackTrace();
      json.put("error", e.getMessage());
      return json;
    }
  }

  public static void postData(String query) {
    try (Connection connect = DriverManager.getConnection(connectionUrl)) {
      connect.createStatement().execute(query);
      connect.close();
    } catch (SQLException e) {
      e.printStackTrace();
    }
  }

  public static void cleanup() {
    try (Connection connection = DriverManager.getConnection(connectionCoreUrl)) {
      Statement statement = connection.createStatement();

      String queryDrop = "DROP DATABASE IF EXISTS pizza_shop_test";
      statement.execute(queryDrop);

      String queryCreate = "CREATE DATABASE pizza_shop_test WITH TEMPLATE pizza_shop OWNER postgres";
      statement.execute(queryCreate);

      connection.close();
      System.out.println("Database cleanup completed.");

    } catch (SQLException e) {
      e.printStackTrace();
    }

  }

}