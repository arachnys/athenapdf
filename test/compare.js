const importTest = (name, path) => {
  describe(name, () => {
    require(path);
  });
};

describe("", () => {
  beforeEach(() => {
    console.log("Running test...");
  });
  importTest("", "./tests/diff");
  importTest("", "./tests/same");
  after(() => {
    console.log("Run tests done");
  });
});
