const { ObjectId, Timestamp } = require("mongodb");

module.exports = {
  async up(db) {
    const collationConfig = {
      collation: {
        caseLevel: true,
        locale: "en_US",
        numericOrdering: true,
        strength: 2,
      },
    };

    await db.createCollection("account", collationConfig);
    await db.collection("counters").insertOne({ name: "account", sid: 0 });

    const ac = await db.collection("account_type").find({}).toArray();

    const initialData = ac.map((v, i) => ({
      account_type: new ObjectId(v._id),
      created_at: new Date(Date.now()),
      isActive: i % 2 === 0 ? true : false,
      isMarkedForDeletion: false,
      isPending: !(i % 2 === 0) ? true : false,
      owner: null,
      sequence_id: null,
      updated_at: new Date(Date.now()),
      user_group_id: null,
    }));

    const r = await db.collection("account").insertMany(initialData);
    console.log(r.insertedIds);

    const xt = await Promise.all(
      Object.values(r.insertedIds).map(async () => {
        const t = await db
          .collection("counters")
          .findOneAndUpdate(
            { name: "account" },
            { $inc: { sid: 1 } },
            { returnDocument: "after" }
          );

        return t;
      })
    );

    await Promise.all(
      Object.values(r.insertedIds).map(async (x, i) => {
        const sorted = xt.sort((y, z) => y.value.sid - z.value.sid);
        await db
          .collection("account")
          .updateOne(
            { _id: new ObjectId(x) },
            { $set: { sequence_id: sorted[i].value.sid } }
          );
      })
    );
  },

  async down(db) {
    await db.collection("account").drop();
    await db.collection("counters").deleteOne({ name: "account" });
  },
};
