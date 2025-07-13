# StardustPilot: Earnium Testnet Automation

Automated wallet management, faucet interaction, and DEX lifecycle management for the Earnium testnet.

StardustPilot is a Go-based automation suite designed specifically for interacting with the Earnium testnet. It provides developers and testers with a robust and reliable tool for managing multiple wallets, acquiring testnet tokens from the faucet, and automating complex DEX operations, including token swaps, liquidity provision, and liquidity removal. The primary goal is to streamline the process of testing and validating smart contracts and decentralized applications on the Earnium platform, significantly reducing manual effort and improving overall efficiency.

This project addresses the common challenges faced when working with testnets, such as the repetitive nature of claiming tokens and the complexity of simulating user interactions on a decentralized exchange. StardustPilot abstracts away these complexities by providing a well-defined API and a configurable execution environment. Users can define custom scripts to automate specific workflows, such as simulating a series of trades, testing different liquidity strategies, or validating the behavior of a smart contract under various conditions. The modular design of the project allows for easy extension and customization to meet the specific needs of different development teams.

The benefits of using StardustPilot are numerous. It significantly reduces the time and effort required to set up and manage testnet environments. It enables developers to quickly iterate on their code and validate their assumptions. It also provides a consistent and reproducible environment for testing, which is crucial for ensuring the quality and reliability of decentralized applications. Furthermore, the project's open-source nature encourages community contributions and ensures that it remains up-to-date with the latest developments in the Earnium ecosystem.

## Key Features

*   **Automated Wallet Management:** Generates and manages multiple Ethereum wallets, storing private keys securely. Uses BIP-39 mnemonics for seed phrase derivation.
*   **Faucet Interaction:** Automatically claims testnet tokens from the Earnium faucet, handling rate limits and retry logic. Implements anti-bot measures evasion.
*   **DEX Swap Automation:** Executes token swaps on the Earnium DEX using predefined parameters, including token pairs, amounts, and slippage tolerance. Supports custom swap paths.
*   **Liquidity Pool Management:** Adds and removes liquidity from Earnium DEX liquidity pools, specifying token amounts and pool IDs. Calculates optimal liquidity ratios.
*   **Scriptable Workflows:** Allows users to define custom scripts to automate complex interactions with the Earnium testnet. Uses a simple domain-specific language (DSL) for script definition.
*   **Transaction Monitoring:** Monitors pending transactions and reports their status, including confirmation and revert events. Provides detailed transaction logs.
*   **Gas Fee Optimization:** Estimates and sets appropriate gas fees to ensure timely transaction execution, considering network congestion and priority.

## Technology Stack

*   **Go (Golang):** The primary programming language used for building the application, chosen for its performance, concurrency, and strong standard library.
*   **Ethereum Go Library (go-ethereum):** Used for interacting with the Ethereum blockchain, including wallet management, transaction signing, and smart contract interaction.
*   **Web3.js (through Go bindings):** Used to facilitate communication with the Earnium testnet via JavaScript bindings exposed to the Go environment.
*   **YAML:** Used for configuration files, providing a human-readable and easily editable format for defining settings and parameters.
*   **GORM (optional):** For database interaction (if persistence of wallet data is required).

## Installation

1.  **Install Go:** Ensure that Go (version 1.18 or later) is installed and configured correctly on your system. You can download it from [https://go.dev/dl/](https://go.dev/dl/).

2.  **Clone the Repository:** Clone the StardustPilot repository from GitHub:
    `git clone https://github.com/Nikeran22/StardustPilot.git`

3.  **Navigate to the Project Directory:**
    `cd StardustPilot`

4.  **Install Dependencies:** Use the Go modules system to install the required dependencies:
    `go mod tidy`

5.  **Build the Application:** Compile the Go code to create an executable binary:
    `go build -o stardustpilot`

## Configuration

StardustPilot requires a configuration file (e.g., `config.yaml`) to define various settings. Here's an example configuration:



*   `rpc_endpoint`: The URL of the Earnium testnet RPC endpoint.
*   `faucet_address`: The address of the faucet smart contract.
*   `private_key`: Your private key (use with caution and only for initial setup).
*   `number_of_wallets`: The number of wallets to generate.
*   `gas_price`: The gas price to use for transactions (in Wei).

## Usage

To run StardustPilot, execute the compiled binary with the configuration file as an argument:

`./stardustpilot --config config.yaml`

Example script for swapping tokens:

`swap.pilot`:


Execute the script with:

`./stardustpilot --config config.yaml --script swap.pilot`

For API documentation, refer to the internal Go documentation within the source code. Use `go doc` or a similar tool to explore the available functions and interfaces.

## Contributing

We welcome contributions to StardustPilot! Please follow these guidelines:

*   Fork the repository and create a branch for your feature or bug fix.
*   Write clear and concise commit messages.
*   Submit a pull request with a detailed description of your changes.
*   Ensure that your code is well-tested and follows the project's coding style.

## License

This project is licensed under the MIT License. See the [LICENSE](https://github.com/Nikeran22/StardustPilot/blob/main/LICENSE) file for details.

## Acknowledgements

We would like to thank the Earnium team for providing the testnet environment and the Ethereum community for their valuable open-source contributions.