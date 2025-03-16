const express = require("express");
const jwt = require("jsonwebtoken");
const User = require("../models/User");
const router = express.Router();


router.post("/register", async (req, res) => {
    try {
      const { username, password } = req.body;
      const newUser = new User({ username, password }); 
      await newUser.save();
      res.status(201).json({ message: "User registered successfully (insecurely)" });
    } catch (error) {
      res.status(500).json({ error: "Internal Server Error" });
    }
  });


  module.exports = router;