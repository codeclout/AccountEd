const assert = require("assert");

module.exports = {
  async up(db, client) {
    await db.createCollection("account_type", {
      collation: {
        caseLevel: true,
        locale: "en_US",
        numericOrdering: true,
        strength: 2,
      },
    });

    const initialData = ["homeschooler", "organization", "study-group"];
    const records = initialData.map((v) => ({
      account_type: v,
    }));

    await db.collection("account_type").insertMany(records);
    const count = await db.collection("account_type").countDocuments();

    assert.strictEqual(records.length, count);
  },

  async down(db, client) {
    await db.collection("account_type").drop();
  },
};
