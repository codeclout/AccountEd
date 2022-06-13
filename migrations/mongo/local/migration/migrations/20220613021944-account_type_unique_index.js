module.exports = {
  async up(db, client) {
    await db
      .collection("account_type")
      .createIndex({ account_type: 1 }, { unique: true });
  },

  async down(db, client) {
    await db.collection("account_type").dropIndex({ account_type: 1 });
  },
};
