<template>
  <div>
    <h1>Notarize PDF</h1>
    <div class="dropbox">
      <input type="file" @change="filesChange($event.target.files);" class="input-file">
      Drag your file(s) here to begin<br> or click to browse<br>
      <div v-if="hash !== null">SHA256 Hash: <b>{{ hash }}</b></div>
      <div v-if="fileName !== null">Name: <b>{{ fileName }}</b></div>
      <div v-if="timestamp !== 0">
        <span id="check">&#10003;</span> VERIFIED at timestamp {{ timestamp }}s with account {{ signer.getAddress() }}
      </div>
    </div>
    <button @click="store()" :disabled="isDisabled">Notarize</button>
    status: <span v-html="status"></span>
  </div>
</template>

<script>
const utils = require('./utils.js')
import {ethers, providers} from "ethers";

const abi = [
  {
    "inputs": [
      {
        "internalType": "bytes32",
        "name": "hash",
        "type": "bytes32"
      }
    ],
    "name": "store",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "address",
        "name": "recipient",
        "type": "address"
      },
      {
        "internalType": "bytes32",
        "name": "hash",
        "type": "bytes32"
      }
    ],
    "name": "verify",
    "outputs": [
      {
        "internalType": "uint256",
        "name": "",
        "type": "uint256"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  }
];
const contractAddress = "0xc8adb8c77068315893d3066634e08706b3d5a3d4";
window.ethereum.enable();
const provider = new providers.Web3Provider(window.ethereum);
const signer = provider.getSigner();
const contract_ro = new ethers.Contract(contractAddress, abi, provider);
const contract_rw = new ethers.Contract(contractAddress, abi, signer);

export default {
  data: function () {
    return {
      hash: null,
      status: 'nothing uploaded',
      isDisabled: true,
      fileName: null,
      timestamp: 0
    }
  },
  created: function () {
    window.addEventListener('load', async () => {
      // Checking if Web3 has been injected by the browser (Mist/MetaMask)
      if (typeof contract_ro === 'undefined') {
        this.status = 'No web3? You should consider trying MetaMask!'
      }
    })
  },
  methods: {
    filesChange(fileList) {
      const fr = new FileReader();
      fr.onload = async (e) => {
        this.hash = utils.stringHash(e.target.result);
        contract_ro.verify(signer.getAddress(), this.hash).then(async (result, error) => {
          if (!error) {
            if (result == 0) {
              this.isDisabled = false
              this.timestamp = 0
              const addr = await signer.getAddress();
              this.status = 'not yet stored from accout: ' + addr
            } else {
              this.isDisabled = true
              this.status = '<b>VERIFIED</b> in the blockchain! Timestamp: ' + result.toNumber()
            }
          } else {
            this.timestamp = 0
            this.isDisabled = false
            this.status = 'error: ' + error
          }
        })
      }
      fr.readAsArrayBuffer(fileList[0]);
      this.fileName = fileList[0].name
    },
    async store() {
      contract_rw.store(this.hash).then((result, error) => {
        if (!error) {
          this.status = 'stored, tx is: ' + result.hash
        } else {
          this.status = 'error: ' + error;
        }
      });
      this.timestamp = 0
      this.isDisabled = true
    }
  }
}
</script>

<style>
.dropbox {
  outline: 2px dashed grey; /* the dash box */
  outline-offset: -10px;
  background: lightcyan;
  color: dimgray;
  padding: 10px 10px;
  min-height: 200px; /* minimum height */
  position: relative;
  cursor: pointer;
}

.input-file {
  opacity: 0; /* invisible but it's there! */
  width: 100%;
  height: 200px;
  position: absolute;
  cursor: pointer;
}

.dropbox:hover {
  background: lightblue; /* when mouse over to the drop zone, change color */
}

.dropbox p {
  font-size: 1.2em;
  text-align: center;
  padding: 50px 0;
}

#check {
  content: "\2713";
  color: green;
  font-size: 2em;
}
</style>
