import React, { useEffect, useState } from "react";
import Form from "react-bootstrap/Form";
import Button from "react-bootstrap/Button";
import "../css/Signup.css";
import axios from "axios";
import { Link, useHistory, useLocation } from "react-router-dom";

export default function Edit() {
  const [phoneNumber, setPhoneNumber] = useState("");
  const [password, setPassword] = useState("");
  const [name, setName] = useState(""); // restaurant name
  const [district, setDistrict] = useState("");
  const [address, setAddress] = useState("");

  const [id, setId] = useState(0);

  let history = useHistory();
  const location = useLocation();

  function handleSubmit(event) {
    event.preventDefault();

    const info = {
      mobile: phoneNumber,
      password: password,
      id: location.state.customer_id,
      name: name,
      district: district,
      address: address,
    };

    axios.post("http://localhost:1323/updatecustomerinfo", info).then(
      (response) => {
        // history.push("/profile/edit");
        history.push({
          pathname: "/profile/edit",
          state: {
            customer_id: location.state.customer_id,
          },
        });
      },
      (error) => {
        console.log(error);
      }
    );
  }

  useEffect(() => {
    setId(location.state.customer_id);

    const info = { id: location.state.customer_id, balance: 0 };
    axios
      .post("http://localhost:1323/getcustomerinfo", info)
      .then((response) => {
        setPhoneNumber(response.data.mobile);
        setPassword(response.data.password);
        setName(response.data.name);
        setDistrict(response.data.district);
        setAddress(response.data.address);
      });
  }, []);

  return (
    <div className="Signup">
      <Form
        onSubmit={handleSubmit}
        style={{ backgroundColor: "#EA738DFF", padding: "20px" }}
      >
        <Form.Group size="lg" controlId="phoneNumber">
          <Form.Label>Phone Number: </Form.Label>
          <Form.Control
            autoFocus
            type="text"
            value={phoneNumber}
            onChange={(e) => setPhoneNumber(e.target.value)}
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
          <Form.Label>Your Name: </Form.Label>
          <Form.Control
            type="text"
            value={name}
            onChange={(e) => setName(e.target.value)}
            style={{ display: "block", margin: "2px" }}
          />
        </Form.Group>

        <Form.Group size="lg" controlId="district">
          <Form.Label>Your District: </Form.Label>
          <Form.Control
            type="text"
            value={district}
            onChange={(e) => setDistrict(e.target.value)}
            style={{ display: "block", margin: "2px" }}
          />
        </Form.Group>

        <Form.Group size="lg" controlId="address">
          <Form.Label>Your Address: </Form.Label>
          <Form.Control
            type="text"
            value={address}
            onChange={(e) => setAddress(e.target.value)}
            style={{ display: "block", margin: "2px" }}
          />
        </Form.Group>

        <Button block size="lg" type="submit">
          Change Information
        </Button>
      </Form>

      <Link
        className="profileLink"
        style={{ textAlign: "center" }}
        to={{
          pathname: "/profile",
          state: {
            customer_id: location.state.customer_id,
          },
        }}
      >
        Back to Profile
      </Link>
    </div>
  );
}
