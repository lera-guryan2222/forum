// import React, { useState } from "react";
// import { useNavigate } from "react-router-dom";
// import { useAuth } from "./AuthContext";
// import styled from "styled-components";

// const Container = styled.div`
//   max-width: 400px;
//   margin: 40px auto;
//   padding: 20px;
//   background: white;
//   border-radius: 8px;
//   box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
// `;

// const Title = styled.h2`
//   text-align: center;
//   color: #333;
//   margin-bottom: 24px;
// `;

// const ErrorMessage = styled.div`
//   color: #dc3545;
//   background-color: #f8d7da;
//   padding: 10px;
//   border-radius: 4px;
//   margin-bottom: 16px;
//   text-align: center;
// `;

// const FormGroup = styled.div`
//   margin-bottom: 16px;
// `;

// const Label = styled.label`
//   display: block;
//   margin-bottom: 8px;
//   color: #333;
// `;

// const Input = styled.input`
//   width: 100%;
//   padding: 10px;
//   border: 1px solid #ddd;
//   border-radius: 4px;
//   font-size: 16px;
//   &:focus {
//     outline: none;
//     border-color: #007bff;
//     box-shadow: 0 0 0 2px rgba(0, 123, 255, 0.25);
//   }
// `;

// const Button = styled.button`
//   width: 100%;
//   padding: 12px;
//   background: #007bff;
//   color: white;
//   border: none;
//   border-radius: 4px;
//   font-size: 16px;
//   cursor: pointer;
//   transition: background-color 0.2s;

//   &:hover {
//     background: #0056b3;
//   }

//   &:disabled {
//     background: #ccc;
//     cursor: not-allowed;
//   }
// `;

// const Register = () => {
//   const [username, setUsername] = useState("");
//   const [email, setEmail] = useState("");
//   const [password, setPassword] = useState("");
//   const [error, setError] = useState("");
//   const [isLoading, setIsLoading] = useState(false);
//   const navigate = useNavigate();
//   const { register } = useAuth();

//   const handleSubmit = async (e) => {
//     e.preventDefault();
//     setIsLoading(true);
//     setError("");

//     try {
//       const result = await register(username, email, password);
//       if (result.success) {
//         navigate("/posts");
//       } else {
//         setError(result.error);
//       }
//     } catch (err) {
//       setError("An unexpected error occurred. Please try again.");
//     } finally {
//       setIsLoading(false);
//     }
//   };

//   return (
//     <Container>
//       <Title>Register</Title>
//       {error && <ErrorMessage>{error}</ErrorMessage>}
//       <form onSubmit={handleSubmit}>
//         <FormGroup>
//           <Label>Username:</Label>
//           <Input
//             type="text"
//             value={username}
//             onChange={(e) => setUsername(e.target.value)}
//             required
//             disabled={isLoading}
//           />
//         </FormGroup>
//         <FormGroup>
//           <Label>Email:</Label>
//           <Input
//             type="email"
//             value={email}
//             onChange={(e) => setEmail(e.target.value)}
//             required
//             disabled={isLoading}
//           />
//         </FormGroup>
//         <FormGroup>
//           <Label>Password:</Label>
//           <Input
//             type="password"
//             value={password}
//             onChange={(e) => setPassword(e.target.value)}
//             required
//             disabled={isLoading}
//             minLength="6"
//           />
//         </FormGroup>
//         <Button type="submit" disabled={isLoading}>
//           {isLoading ? "Registering..." : "Register"}
//         </Button>
//       </form>
//     </Container>
//   );
// };

// export default Register;
import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import { useAuth } from "./AuthContext";
import styled from "styled-components";

const Container = styled.div`
  max-width: 400px;
  margin: 40px auto;
  padding: 20px;
  background: white;
  border-radius: 8px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
`;

const Title = styled.h2`
  text-align: center;
  color: #333;
  margin-bottom: 24px;
`;

const ErrorMessage = styled.div`
  color: #dc3545;
  background-color: #f8d7da;
  padding: 10px;
  border-radius: 4px;
  margin-bottom: 16px;
  text-align: center;
`;

const FormGroup = styled.div`
  margin-bottom: 16px;
`;

const Label = styled.label`
  display: block;
  margin-bottom: 8px;
  color: #333;
`;

const Input = styled.input`
  width: 100%;
  padding: 10px;
  border: 1px solid #ddd;
  border-radius: 4px;
  font-size: 16px;
  &:focus {
    outline: none;
    border-color: #007bff;
    box-shadow: 0 0 0 2px rgba(0, 123, 255, 0.25);
  }
`;

const Button = styled.button`
  width: 100%;
  padding: 12px;
  background: #007bff;
  color: white;
  border: none;
  border-radius: 4px;
  font-size: 16px;
  cursor: pointer;
  transition: background-color 0.2s;

  &:hover {
    background: #0056b3;
  }

  &:disabled {
    background: #ccc;
    cursor: not-allowed;
  }
`;

const Register = () => {
  const [username, setUsername] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState("");
  const [isLoading, setIsLoading] = useState(false);
  const navigate = useNavigate();
  const { register } = useAuth();

  const handleSubmit = async (e) => {
  e.preventDefault();
  setIsLoading(true);
  setError("");

  try {
    const result = await register(username, email, password);
    if (result.success) {
      navigate("/posts");
    } else {
      setError(result.error || "Registration failed");
    }
  } catch (err) {
    setError(err.message || "An unexpected error occurred. Please try again.");
  } finally {
    setIsLoading(false);
  }
};

  return (
    <Container>
      <Title>Register</Title>
      {error && <ErrorMessage>{error}</ErrorMessage>}
      <form onSubmit={handleSubmit}>
        <FormGroup>
          <Label>Username:</Label>
          <Input
            type="text"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
            required
            disabled={isLoading}
          />
        </FormGroup>
        <FormGroup>
          <Label>Email:</Label>
          <Input
            type="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            required
            disabled={isLoading}
          />
        </FormGroup>
        <FormGroup>
          <Label>Password:</Label>
          <Input
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            required
            disabled={isLoading}
            minLength="6"
          />
        </FormGroup>
        <Button type="submit" disabled={isLoading}>
          {isLoading ? "Registering..." : "Register"}
        </Button>
      </form>
    </Container>
  );
};

export default Register;
