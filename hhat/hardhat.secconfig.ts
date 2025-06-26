import { HardhatUserConfig } from "hardhat/config";
import "@nomicfoundation/hardhat-toolbox";

const config: HardhatUserConfig = {
  solidity: "0.8.28", 
  networks: {
    hardhat: {
      chainId: 31338,
      accounts: [
        {
            privateKey: "0x6b69db447fce3fc716ba7bdf5d40014dc10f84de305aa3cb0ee443a89c1c9db5",
            balance: "10000000000000000000000",
        },
        {
            privateKey: "0xa579f302bdbeaa73098c38e89be77050b8139d4b94cfae2c5dfb1ae7c6c5a973",
            balance: "10000000000000000000000",
        },
        {
            privateKey: "0x201be0fc331eb0ff17b1fe2ce26b1edc393c1f47e09c5ae8e352944a12f23503",
            balance: "10000000000000000000000",
        },
        {
            privateKey: "0xa6ebf9461a5a2a5c1cf89351a98ad2a14b3fe4f5e750edcf3a8bc3363f8c0aa3",
            balance: "10000000000000000000000",
        },
        {
            privateKey: "0x3fa406e9946f5241a226f7f5f2f37d734f54c1bfb3bdfb8ab4f21acb2cc7e0c1",
            balance: "10000000000000000000000",
        },
      ]
    },
  },
};

export default config;