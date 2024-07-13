import { ChangeEvent, useState } from "react";
import GLogo from "/G.svg";
import { BrowserProvider } from "ethers";
import { isConnected, requestAccess } from "@stellar/freighter-api";
import Swal from "sweetalert2";
import withReactContent from 'sweetalert2-react-content';

function App() {
  const [formData, setFormData] = useState({
    username: "",
    password: "",
    address: "",
    network: "",
  });

  const handleChange = (e: ChangeEvent<HTMLInputElement>) => {
    setFormData({
      ...formData,
      [e.target.name]: e.target.value,
    });
  };

  const connectMetaMask = async () => {
    if (typeof window.ethereum === 'undefined') {
      Swal.fire({
        title: "MetaMask is not installed",
        text: "Please install MetaMask to use this feature.",
        icon: "error",
      });
      return;
    }
    else {
    const provider = new BrowserProvider(window.ethereum);
    const signer = await provider.getSigner();
    setFormData({
      ...formData,
      address: signer.address,
      network: "ethereum",
    });
  }
  };

  const connectFreighter = async () => {
    if (typeof window.ethereum === 'undefined') {
      Swal.fire({
        title: "Freighter wallet is not installed",
        text: "Please install Freighter wallet to use this feature.",
        icon: "error",
      });
      return;
    }
    else if (await isConnected()) {
      const publicKey = await requestAccess();
      setFormData({
        ...formData,
        address: publicKey,
        network: "stellar",
      });
    }
  };
  const MySwal = withReactContent(Swal);
  const handleSubmit = async (e: ChangeEvent<HTMLFormElement>) => {
    e.preventDefault();
    MySwal.fire({
      title: <p>Loading</p>,
      didOpen: () => {
        // `MySwal` is a subclass of `Swal` with all the same instance & static methods
        MySwal.showLoading()
      },
    })
    try {
      // console.log(formData);
      const response = await fetch("register", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: "Bearer token",
        },
        body: JSON.stringify(formData),
      });
      MySwal.close(); 
      if (!response.ok) {
        Swal.fire({
          title: "Registration Failure",
          text: "Try again later",
          icon: "error",
        });
      }

      switch (response.status) {
        case 201:
          const data = await response.json();
          console.log("Success:", data);
          Swal.fire({
            title: "Registration Success",
            text: "Download & Battle",
            showConfirmButton: true,
            icon: "success",
            confirmButtonText: "Download",
          }).then((result) => {
            if (result.isConfirmed) {
              Swal.fire("Complete!", "", "success");
              setFormData({
                username: "",
                password: "",
                address: "",
                network: "",
              })
            }
          });
          break;
        case 400:
          Swal.fire({
            title: "Duplication Detected",
            text: "Make sure the username & address are unique.",
            icon: "warning",
          });
          break;
        case 430:
          Swal.fire({
            title: "Validation Error",
            text: "Make sure fields are not empty.",
            icon: "warning",
          });
          break;
        default:
          Swal.fire({
            title: "Unwarranted Response",
            text: "Unable to process.",
            icon: "error",
          });
          break;
      }
    } catch (error) {
      console.error("Error:", error);
      Swal.fire({
        title: "Registration Failure",
        text: "Kindly try again.",
        icon: "error",
      });
    }
  };

  return (
    <>
      <div className="flex min-h-full flex-1 flex-col justify-center px-6 py-12 lg:px-8">
        <div className="sm:mx-auto sm:w-full sm:max-w-sm">
          <img alt="Your Company" src={GLogo} className="mx-auto h-12 w-auto" />
          <h2 className="mt-2 text-center text-2xl font-bold leading-9 tracking-tight text-gray-900">
            GRIMOIRE
          </h2>
          <h3 className="text-center text-xl">Registration</h3>
        </div>

        <div className="mt-5 sm:mx-auto sm:w-full sm:max-w-sm">
          <form className="space-y-6" onSubmit={handleSubmit}>
            <div>
              <label
                htmlFor="username"
                className="block text-sm font-medium leading-6 text-gray-900"
              >
                Username
              </label>
              <div className="mt-2">
                <input
                  id="username"
                  name="username"
                  type="text"
                  maxLength={20}
                  required
                  value={formData.username}
                  onChange={handleChange}
                  className="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                />
              </div>
            </div>

            <div>
              <div className="flex items-center justify-between">
                <label
                  htmlFor="password"
                  className="block text-sm font-medium leading-6 text-gray-900"
                >
                  Password
                </label>
              </div>
              <div className="mt-2">
                <input
                  id="password"
                  name="password"
                  type="password"
                  required
                  maxLength={50}
                  value={formData.password}
                  onChange={handleChange}
                  className="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                />
              </div>
            </div>

            <div className="justify-center">
              <button
                type="button"
                className="text-white w-full bg-gradient-to-br from-purple-600 to-blue-500 hover:bg-gradient-to-bl focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center me-2 mb-2"
                onClick={connectFreighter}
              >
                Connect Freighter
              </button>
              <span className="px-2 py-2 flex justify-center">or</span>
              <button
                type="button"
                className="text-white w-full bg-gradient-to-br from-pink-500 to-orange-400 hover:bg-gradient-to-bl focus:ring-4 focus:outline-none focus:ring-pink-200 font-medium rounded-lg text-sm px-5 py-2.5 text-center me-2 mb-2"
                onClick={connectMetaMask}
              >
                Connect MetaMask
              </button>
            </div>

            <div>
              <div className="mt-2">
                <input
                  id="address"
                  name="address"
                  type="text"
                  required
                  readOnly
                  value={formData.address}
                  className="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                />
              </div>
            </div>

            <div>
              <button
                type="submit"
                className="flex w-full justify-center rounded-md bg-indigo-600 px-3 py-1.5 text-sm font-semibold leading-6 text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
              >
                Register
              </button>
            </div>
          </form>
        </div>
      </div>
    </>
  );
}

export default App;
