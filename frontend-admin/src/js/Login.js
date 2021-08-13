import React, { useState } from "react";
import Form from "react-bootstrap/Form";
import Button from "react-bootstrap/Button";
import "../css/Login.css";
import axios from "axios";
import { Link, useHistory } from "react-router-dom";

export default function Login() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  function validateForm() {
    return email.length > 0 && password.length > 0;
  }

  let history = useHistory();

  function handleSubmit(event) {
    event.preventDefault();
    const info = { email: email, password: password };
    axios.post("http://localhost:1323/loginadmin", info).then(
      (response) => {
        // history.push("/profile");
        history.push({
          pathname: "/profile",
          state: {
            restaurant_id: response.data.restaurant_id,
          },
        });
      },
      (error) => {
        console.log(error);
      }
    );
  }

  return (
    <div className="Login">
      <Form
        onSubmit={handleSubmit}
        style={{ backgroundColor: "#EA738DFF", padding: "20px" }}
      >
        <Form.Group size="lg" controlId="email">
          <Form.Label>Email: </Form.Label>
          <Form.Control
            autoFocus
            placeholder="Enter email"
            type="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            style={{ display: "block", margin: "2px" }}
          />
        </Form.Group>
        <Form.Group size="lg" controlId="password">
          <Form.Label>Password: </Form.Label>
          <Form.Control
            type="password"
            placeholder="Enter password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            style={{ display: "block", margin: "2px" }}
          />
        </Form.Group>
        <Button block size="lg" type="submit" disabled={!validateForm()}>
          Login
        </Button>
      </Form>

      <div className="signupLink">
        <Link to="/signup">Signup</Link>
      </div>
    </div>
  );
}
