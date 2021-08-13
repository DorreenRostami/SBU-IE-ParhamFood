import React, { useEffect, useState } from "react";
import Form from "react-bootstrap/Form";
import Button from "react-bootstrap/Button";
import "../css/Signup.css";
import axios from "axios";
import { Link, useHistory, useLocation } from "react-router-dom";

export default function Edit() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [name, setName] = useState(""); // restaurant name
  const [district, setDistrict] = useState("");
  const [address, setAddress] = useState("");
  const [open, setOpen] = useState(0);
  const [close, setClose] = useState(0);
  const [fixedCost, setFixedCost] = useState(0);
  const [fixedTime, setFixedTime] = useState(0);

  const [id, setId] = useState(0);

  let history = useHistory();
  const location = useLocation();

  function handleSubmit(event) {
    event.preventDefault();

    const resInfo = {
      email: email,
      password: password,
      name: name,
      district: district,
      address: address,
      open: parseInt(open, 10),
      close: parseInt(close, 10),
      fixed_cost: parseInt(fixedCost, 10),
      fixed_minute: parseInt(fixedTime, 10),
    };

    const info = {
      restaurant_id: id,
      restaurant_info: resInfo,
    };

    axios.post("http://localhost:1323/updaterestaurantinfo", info).then(
      (response) => {
        history.push({
          pathname: "/profile/edit",
          state: {
            restaurant_id: id,
          },
        });
      },
      (error) => {
        console.log(error);
      }
    );
  }

  useEffect(() => {
    setId(location.state.restaurant_id);
    const info = { restaurant_id: location.state.restaurant_id };
    axios
      .post("http://localhost:1323/getrestaurantinfo", info)
      .then((response) => {
        setEmail(response.data.email);
        setPassword(response.data.password);
        setName(response.data.name);
        setDistrict(response.data.district);
        setAddress(response.data.address);
        setOpen(response.data.open);
        setClose(response.data.close);
        setFixedCost(response.data.fixed_cost);
        setFixedTime(response.data.fixed_minute);
      });
  }, []);

  return (
    <div className="Signup">
      <Form
        onSubmit={handleSubmit}
        style={{ backgroundColor: "#EA738DFF", padding: "20px" }}
      >
        <Form.Group size="lg" controlId="email">
          <Form.Label>Email: </Form.Label>
          <Form.Control
            autoFocus
            placeholder="name@example.com"
            type="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            style={{ display: "block", margin: "2px" }}
          />
        </Form.Group>

        <Form.Group size="lg" controlId="password">
          <Form.Label>Password: </Form.Label>
          <Form.Control
            type="text"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            style={{ display: "block", margin: "2px" }}
          />
        </Form.Group>

        <Form.Group size="lg" controlId="name">
          <Form.Label>Restaurant Name: </Form.Label>
          <Form.Control
            type="text"
            value={name}
            onChange={(e) => setName(e.target.value)}
            style={{ display: "block", margin: "2px" }}
          />
        </Form.Group>

        <Form.Group size="lg" controlId="district">
          <Form.Label>Restaurant District: </Form.Label>
          <Form.Control
            type="text"
            value={district}
            onChange={(e) => setDistrict(e.target.value)}
            style={{ display: "block", margin: "2px" }}
          />
        </Form.Group>

        <Form.Group size="lg" controlId="address">
          <Form.Label>Restaurant Address: </Form.Label>
          <Form.Control
            type="text"
            value={address}
            onChange={(e) => setAddress(e.target.value)}
            style={{ display: "block", margin: "2px" }}
          />
        </Form.Group>

        <Form.Group size="lg" controlId="open">
          <Form.Label>Restaurant Open Time: </Form.Label>
          <Form.Control
            type="number"
            value={open}
            onChange={(e) => setOpen(e.target.value)}
            style={{ display: "block", margin: "2px" }}
          />
        </Form.Group>

        <Form.Group size="lg" controlId="close">
          <Form.Label>Restaurant Close Time: </Form.Label>
          <Form.Control
            type="number"
            value={close}
            onChange={(e) => setClose(e.target.value)}
            style={{ display: "block", margin: "2px" }}
          />
        </Form.Group>

        <Form.Group size="lg" controlId="fixedCost">
          <Form.Label>Restaurant Fixed Cost: </Form.Label>
          <Form.Control
            type="number"
            value={fixedCost}
            onChange={(e) => setFixedCost(e.target.value)}
            style={{ display: "block", margin: "2px" }}
          />
        </Form.Group>

        <Form.Group size="lg" controlId="fixedTime">
          <Form.Label>Restaurant Fixed Time: </Form.Label>
          <Form.Control
            type="number"
            value={fixedTime}
            onChange={(e) => setFixedTime(e.target.value)}
            style={{ display: "block", margin: "2px" }}
          />
        </Form.Group>

        <Button block size="lg" type="submit">
          Change Information
        </Button>
      </Form>

      <div className="loginLink" style={{ textAlign: "center" }}>
        <Link
          to={{
            pathname: "/profile",
            state: {
              restaurant_id: id,
            },
          }}
          style={{ margin: "4rem" }}
        >
          Back to Profile
        </Link>
      </div>
    </div>
  );
}
