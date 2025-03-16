const request = require("supertest");
const app = require("../server");
const mongoose = require("mongoose");
const bcrypt = require("bcrypt");
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

describe("User Registration API", () => {
  it("should register a user with a hashed password", async () => {
    const res = await request(app)
      .post("/register") // Fixed endpoint
      .send({ username: "testuser", password: "plaintextpassword" });

    expect(res.status).toBe(201);
    expect(res.body).toHaveProperty("message");

    const user = await User.findOne({ username: "testuser" });
    expect(user).not.toBeNull();
    
    // Check if the stored password is hashed
    const isPasswordHashed = await bcrypt.compare("plaintextpassword", user.password);
    expect(isPasswordHashed).toBe(true);
  });

  it("should not allow duplicate usernames", async () => {
    await request(app).post("/register").send({ username: "duplicateUser", password: "password123" });

    const res = await request(app).post("/register").send({ username: "duplicateUser", password: "newpassword" });

    expect(res.status).toBe(400);
    expect(res.body).toHaveProperty("error", "Username already exists");
  });
});
