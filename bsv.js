const bsv = require('bsv');
const readline = require('readline');

// Create a readline interface for user input
const rl = readline.createInterface({
  input: process.stdin,
  output: process.stdout
});

// Function to get user input
function getInput(prompt) {
  return new Promise((resolve) => {
    rl.question(prompt, (answer) => {
      resolve(answer);
    });
  });
}

// Function to generate the transaction hex
function generateTransactionHex(lockingScript, unlockingScript) {
  // Create a new transaction
  const transaction = new bsv.Transaction();

  // Set the locking script details
  const lockingScriptObj = new bsv.Script(lockingScript);

  // Add inputs to the transaction
  transaction.addInput(
    new bsv.Transaction.Input.PublicKeyHash({
      output: new bsv.Transaction.Output({
        script: lockingScriptObj,
        satoshis: 1000000 // Set the amount to spend
      }),
      prevTxId: '140b8f555b73d91c6bd56b39819df9d08b9fdbd5da06b67f3e0597c4f53f24fc', // Set the previous transaction ID
      outputIndex: 0 // Set the output index of the previous transaction
    })
  );

  // Add outputs to the transaction
  transaction.addOutput(
    new bsv.Transaction.Output({
      script: bsv.Script.buildPublicKeyHashOut('1BvBMSEYstWetqTFn5Au4m4GFg7xJaNVN2'),
      satoshis: 900000 // Set the amount to send
    })
  );

  // Set the unlocking script
  const unlockingScriptObj = new bsv.Script(unlockingScript);
  transaction.inputs[0].setScript(unlockingScriptObj);

  // Return the transaction hex
  return transaction.serialize();
}

// Run the code
(async () => {
  // Prompt the user for the locking script and unlocking script
  const lockingScript = await getInput('Enter the locking script: ');
  const unlockingScript = await getInput('Enter the unlocking script: ');

  // Generate the transaction hex
  const transactionHex = generateTransactionHex(lockingScript, unlockingScript);
  console.log('Transaction hex:', transactionHex);

  // Close the readline interface
  rl.close();
})();
