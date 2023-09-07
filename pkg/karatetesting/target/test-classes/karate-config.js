function fn() {
  var env = karate.env;
  karate.log("karate.env system property was:", env);
  if (!env) {
    env = "dev";
  }

  let config = {
    env: env,
    myVarName: "someValue",
    refreshToken:
      "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJjdXN0b21fa2V5IjoiZjU0NWJkNmE1OTkwZDg3NjY0MzdhZDUwOWQyMWU2NGRjZTUwOGVkOTA5NTZjMmM3NGE0YjA2MmQzOTE3YTAzMyIsImtleV90eXBlIjoicmVmcmVzaCIsImlzcyI6ImJvb2tpdGUuYXV0aC5zZXJ2aWNlIn0.G6NxB9LGGszVoUXNpnZGadrBMbbhjRsAFGcTI5TINnDErTbfmsC-bqGuD2yvE7sr0NLtj4si5gtg8TfUxYekX5A_Te334p-Q4NyrWOArS_i48-CHKiA9G76gRVu82J0XpeMULr0crUYHBAtK2L1mYbQiQh1aXmnYw_KfYcA8dr5C6UJiI-aJLLycJWmzIUN_aFe3DI_u2ce0P-7fjfb5dNAgKrV3hZkbPLFJur0SVN3G-CGRgBFh9g7CKUvOTJkD13I0L3jvPzgWTEf7O_11LfYE3hIz5ltYc0bCoLd7DmC9lJLgo5gtkUbDbzwrIcDHlOztVWRL550LnKUCf6m5Ig",
    url: "http://194.233.85.132:9000/api/v1/refresh-token",
  };

  let tokenFetch = Java.type("helpers.TokenFetch");
  let v = tokenFetch.getAccessToken(config.url, config.refreshToken);

  config.accessToken = v.data.access_token;
  
  return config;
}
