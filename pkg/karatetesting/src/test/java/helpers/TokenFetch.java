package helpers;

import java.io.IOException;
import java.net.URI;
import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;
import java.util.Base64;

import com.oracle.truffle.regex.tregex.util.json.JsonObject;

import net.minidev.json.JSONObject;
import net.minidev.json.JSONValue;

public class TokenFetch {

  public static JSONObject getAccessToken(String url, String refreshToken) throws IOException, InterruptedException {
    HttpClient client = HttpClient.newHttpClient();

    // Create an Authorization header by encoding your refresh token (assuming it's
    // a string)
    String authHeader = "Bearer " + refreshToken;

    HttpRequest request = HttpRequest.newBuilder()
        .uri(URI.create(url))
        .header("Authorization", authHeader) // Add the Authorization header
        .build();

    HttpResponse<String> response = client.send(request,
        HttpResponse.BodyHandlers.ofString());

    // Check if the response status is 200 (OK) before processing the JSON
    if (response.statusCode() == 200) {
      String responseBody = response.body();

      JSONObject jsonResponse = (JSONObject) JSONValue.parse(responseBody);

      return jsonResponse;
    } else {
      System.err.println("Request failed with status code: " + response.statusCode());
    }

    return new JSONObject();
  }
}
