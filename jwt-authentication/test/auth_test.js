const request = require("supertest");
const app = require("../server");
const mongoose = require("mongoose");
const User = require("../models/User");

beforeAll(async () => {
  await mongoose.connect(process.env.TEST_DB_URI, {
    useNewUrlParser: true,
    useUnifiedTopology: true,
  });
});

afterAll(async () => {
  await mongoose.connection.dropDatabase();
  await mongoose.connection.close();
});

describe("User Registration", () => {
  it("should register a user", async () => {
    const res = await request(app)
      .post("/api/auth/register")
      .send({ username: "testuser", password: "plaintextpassword" });

    expect(res.status).toBe(201);
    expect(res.body.message).toBe("User registered successfully (insecurely)");

    const user = await User.findOne({ username: "testuser" });
    expect(user).not.toBeNull();
    expect(user.password).toBe("plaintextpassword");
  });
});