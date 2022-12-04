const { ObjectId, Double } = require("mongodb");

module.exports = {
  async up(db, client) {
    await db.createCollection("membership");
    await db.createCollection("counters", {
      collation: {
        caseLevel: true,
        locale: "en_US",
        numericOrdering: true,
        strength: 2,
      },
    });

    await db.collection("counters").insertOne({ name: "membership", sid: 0 });

    const ac = await db
      .collection("account_type")
      .find({ account_type: "homeschooler" })
      .toArray();
    const homeschoolerAccountTypeId = new ObjectId(ac[0]._id);

    const initialData = [
      {
        account_type: homeschoolerAccountTypeId,
        active: true,
        created_at: new Date(Date.now()),
        expires_at: new Date(Date.now() + 31556952000),
        price: new Double(20.0),
        sequence_id: null,
      },
    ];

    await db.collection("membership").insertMany(initialData);

    const membershipCounter = await db
      .collection("counters")
      .findOneAndUpdate(
        { name: "membership" },
        { $inc: { sid: 1 } },
        { returnDocument: "after" }
      );

    await db
      .collection("membership")
      .updateOne(
        { account_type: homeschoolerAccountTypeId },
        { $set: { sequence_id: membershipCounter.value.sid } }
      );
  },

  async down(db, client) {
    await db.collection("membership").drop();
    await db.collection("counters").drop();
  },
};
