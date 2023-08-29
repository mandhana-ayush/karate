function fn() {
  var env = karate.env; // get system property 'karate.env'
  karate.log("karate.env system property was:", env);
  if (!env) {
    env = "dev";
  }
  var config = {
    env: env,
    myVarName: "someValue",
  };

  karate.configure("afterFeature", function () {
    karate.log("Hey this is after feature....running after every feature.");
  });

  return config;
}
