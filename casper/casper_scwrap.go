// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package casper

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// CasperABI is the input ABI used to generate the binding from.
const CasperABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"addToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"srcIp\",\"type\":\"string\"}],\"name\":\"getPingTarget\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"confirmDownload\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"ip\",\"type\":\"string\"},{\"name\":\"size\",\"type\":\"uint256\"}],\"name\":\"confirmUpload\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"ip\",\"type\":\"string\"},{\"name\":\"size\",\"type\":\"uint256\"}],\"name\":\"notifySpaceFreed\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"ip\",\"type\":\"string\"}],\"name\":\"removeProviderMachine\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"clearPrepay\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"client\",\"type\":\"address\"}],\"name\":\"isPrepaid\",\"outputs\":[{\"name\":\"prepaid\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"size\",\"type\":\"uint256\"}],\"name\":\"getPeers\",\"outputs\":[{\"name\":\"ip\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"ip\",\"type\":\"string\"}],\"name\":\"setBootstrap\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"pay\",\"type\":\"uint256\"}],\"name\":\"prePay\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"confirmTransfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"ip\",\"type\":\"string\"},{\"name\":\"size\",\"type\":\"uint256\"}],\"name\":\"registerProvider\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"uuid\",\"type\":\"uint128\"}],\"name\":\"notifyDelete\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getBootstrap\",\"outputs\":[{\"name\":\"ip\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"ip\",\"type\":\"string\"},{\"name\":\"success\",\"type\":\"bool\"}],\"name\":\"sendPingResult\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"ip\",\"type\":\"string\"}],\"name\":\"machineInformation\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"ip\",\"type\":\"string\"},{\"name\":\"size\",\"type\":\"uint256\"}],\"name\":\"addProviderMachine\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"info\",\"type\":\"string\"}],\"name\":\"Log\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"ProviderCheckEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"val\",\"type\":\"string\"}],\"name\":\"ReturnString\",\"type\":\"event\"}]"

// CasperBin is the compiled bytecode used for deploying new contracts.
const CasperBin = `0x6060604052608060005534156200001557600080fd5b601f620000246004826200002b565b50620000c6565b81548183558181151162000052576000838152602090206200005291810190830162000057565b505050565b6200008391905b808211156200007f57600062000075828262000086565b506001016200005e565b5090565b90565b5080546000825590600052602060002090810190620000a69190620000a9565b50565b6200008391905b808211156200007f5760008155600101620000b0565b61239880620000d66000396000f3006060604052600436106100fb5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663179d375c81146101005780631acaa8331461011857806328748e91146101695780633f65a6d01461017c5780634d2ab7e4146101cf5780635107e4ee146102225780636e2e18871461027357806373e07e14146102865780639b92f9b1146102b9578063a746dadb14610346578063aac80f4714610397578063baa4930114610169578063bb66e09d146103ad578063c72f7c5414610400578063d9290af914610428578063dc3dd5781461043b578063f8c5df3314610490578063fde18f47146104f3575b600080fd5b341561010b57600080fd5b610116600435610546565b005b341561012357600080fd5b61011660046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284375094965061056795505050505050565b341561017457600080fd5b610116610dae565b341561018757600080fd5b61011660046024813581810190830135806020601f820181900481020160405190810160405281815292919060208401838380828437509496505093359350610db092505050565b34156101da57600080fd5b61011660046024813581810190830135806020601f820181900481020160405190810160405281815292919060208401838380828437509496505093359350610ed992505050565b341561022d57600080fd5b61011660046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284375094965061103795505050505050565b341561027e57600080fd5b6101166112eb565b341561029157600080fd5b6102a5600160a060020a0360043516611306565b604051901515815260200160405180910390f35b34156102c457600080fd5b6102cf600435611323565b60405160208082528190810183818151815260200191508051906020019080838360005b8381101561030b5780820151838201526020016102f3565b50505050905090810190601f1680156103385780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561035157600080fd5b61011660046024813581810190830135806020601f820181900481020160405190810160405281815292919060208401838380828437509496506114c395505050505050565b34156103a257600080fd5b6101166004356115c7565b34156103b857600080fd5b61011660046024813581810190830135806020601f82018190048102016040519081016040528181529291906020840183838082843750949650509335935061164e92505050565b341561040b57600080fd5b6101166fffffffffffffffffffffffffffffffff6004351661165c565b341561043357600080fd5b6102cf61165f565b341561044657600080fd5b61011660046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284375094965050505091351515915061174e9050565b341561049b57600080fd5b6104e160046024813581810190830135806020601f82018190048102016040519081016040528181529291906020840183838082843750949650611b1695505050505050565b60405190815260200160405180910390f35b34156104fe57600080fd5b61011660046024813581810190830135806020601f820181900481020160405190810160405281815292919060208401838380828437509496505093359350611b9392505050565b600160a060020a033316600090815260026020526040902080549091019055565b600080600080600080600060036000896040518082805190602001908083835b602083106105a65780518252601f199092019160209182019101610587565b6001836020036101000a0380198251168184511617909252505050919091019250604091505051908190039020815260208101919091526040016000206007810154909750610100900460ff16156105fd57600080fd5b865433600160a060020a0390811691161461061757600080fd5b6007546008546001909101901080156106355750600787015460ff16155b801561064357506005870154155b156106f15760078701805460ff19166001908117909155600880549091810161066c83826121ee565b916000526020600020900160008960010160405180828054600181600116156101000203166002900480156106d85780601f106106b65761010080835404028352918201916106d8565b820191906000526020600020905b8154815290600101906020018083116106c4575b5050915050604051908190039020909155506107e19050565b60075460085460019091019011156107e1576008546107169060009060001901611ed9565b955060006003600060088981548110151561072d57fe5b60009182526020808320919091015483528201929092526040019020600701805460ff1916911515919091179055600880548790811061076957fe5b600091825260208220015560085460019011801561078d5750600854600019018614155b156107cc576008805460001981019081106107a457fe5b9060005260206000209001546008878154811015156107bf57fe5b6000918252602090912001555b6008805460001901906107df90826121ee565b505b6007546000955085901180156107fb5750600787015460ff165b1561091a57600780546108149060009060001901611ed9565b8154811061081e57fe5b90600052602060002090015494507f3ba3d58761f64d3f96a3d0808397c076d4936a34f18b36fb0b297a93ee8a83b960405160405180910390a1600085815260036020526040908190207f9e91c0d67653515860c12ff5c5211b3ef5b98565f70728283708dd7bf754c4f19160019091019051602080825282546002600019610100600184161502019091160490820181905281906040820190849080156109075780601f106108dc57610100808354040283529160200191610907565b820191906000526020600020905b8154815290600101906020018083116108ea57829003601f168201915b50509250505060405180910390a1610da4565b61092660006001611ed9565b9350836001148015610951575060048054601e90811061094257fe5b60009182526020909120015415155b156109b65760048054601e90811061096557fe5b9060005260206000209001610999600060016004601e81548110151561098757fe5b60009182526020909120015403611ed9565b815481106109a357fe5b9060005260206000209001549450610ce6565b6109c26000601d611ed9565b9250829150600090505b6000198213806109dc5750601f83125b15610ce657600019821315610b955760048054839081106109f957fe5b60009182526020909120015415610b8d57610a206000600160048581548110151561098757fe5b9050876040518082805190602001908083835b60208310610a525780518252601f199092019160209182019101610a33565b6001836020036101000a03801982511681845116179092525050509190910192506040915050519081900390206004805484908110610a8d57fe5b906000526020600020900182815481101515610aa557fe5b60009182526020909120015414610adc576004805483908110610ac457fe5b9060005260206000209001818154811015156109a357fe5b6001600483815481101515610aed57fe5b6000918252602090912001541115610b8d576001600483815481101515610b1057fe5b600091825260209091200154038114610b5f576004805483908110610b3157fe5b906000526020600020900181600101815481101515610b4c57fe5b9060005260206000209001549450610b88565b6004805483908110610b6d57fe5b9060005260206000209001600182038154811015156109a357fe5b610ce6565b600019909101905b601f831215610ce1576004805484908110610bac57fe5b60009182526020909120015415610cda57610bd36000600160048681548110151561098757fe5b9050876040518082805190602001908083835b60208310610c055780518252601f199092019160209182019101610be6565b6001836020036101000a03801982511681845116179092525050509190910192506040915050519081900390206004805485908110610c4057fe5b906000526020600020900182815481101515610c5857fe5b60009182526020909120015414610c77576004805484908110610ac457fe5b6001600484815481101515610c8857fe5b6000918252602090912001541115610cda576001600484815481101515610cab57fe5b600091825260209091200154038114610ccc576004805484908110610b3157fe5b6004805484908110610b6d57fe5b6001909201915b6109cc565b600085815260036020526040908190207f9e91c0d67653515860c12ff5c5211b3ef5b98565f70728283708dd7bf754c4f1916001909101905160208082528254600260001961010060018416150201909116049082018190528190604082019084908015610d955780601f10610d6a57610100808354040283529160200191610d95565b820191906000526020600020905b815481529060010190602001808311610d7857829003601f168201915b50509250505060405180910390a15b5050505050505050565b565b600160a060020a033316600090815260016020526040812054819011610dd557600080fd5b60036000846040518082805190602001908083835b60208310610e095780518252601f199092019160209182019101610dea565b6001836020036101000a038019825116818451161790925250505091909101925060409150505190819003902081526020810191909152604001600020600a81015490915082901015610e5b57600080fd5b600a810180548390039055610e6f83611f26565b7fcf34ef537ac33ee1ac626ca1587a0a7e8e51561e5514f8cb36afa1c5102b3bab60405160208082526017908201527f5375636365737366756c6c7920636f6e6669726d6564210000000000000000006040808301919091526060909101905180910390a1505050565b600060036000846040518082805190602001908083835b60208310610f0f5780518252601f199092019160209182019101610ef0565b6001836020036101000a03801982511681845116179092525050509190910192506040915050519081900390208152602081019190915260400160002060098101549091501515610f5f57600080fd5b600a8101805483019081905582901015610f7857600080fd5b600a81015460098201541015610f8d57600080fd5b611032816001018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156110285780601f10610ffd57610100808354040283529160200191611028565b820191906000526020600020905b81548152906001019060200180831161100b57829003601f168201915b5050505050611f26565b505050565b60008060036000846040518082805190602001908083835b6020831061106e5780518252601f19909201916020918201910161104f565b6001836020036101000a038019825116818451161790925250505091909101925060409150505190819003902081526020810191909152604001600020600981015490925015156110be57600080fd5b600482600201548154811015156110d157fe5b906000526020600020900182600301548154811015156110ed57fe5b6000918252602082200155600282015460048054909190811061110c57fe5b6000918252602090912001549050600181118015611131575060018103826003015414155b156111ac576004826002015481548110151561114957fe5b90600052602060002090016001820381548110151561116457fe5b9060005260206000209001546004836002015481548110151561118357fe5b9060005260206000209001836003015481548110151561119f57fe5b6000918252602090912001555b60018103600483600201548154811015156111c357fe5b9060005260206000209001816111d991906121ee565b5060098201548254600160a060020a03166000908152600260205260408082206001018054939093039092556003918590518082805190602001908083835b602083106112375780518252601f199092019160209182019101611218565b6001836020036101000a0380198251168184511617909252505050919091019250604091505051908190039020815260208101919091526040016000908120805473ffffffffffffffffffffffffffffffffffffffff191681559061129f6001830182612212565b506000600282018190556003820181905560048201819055600582018190556006820181905560078201805461ffff191690556008820181905560098201819055600a90910155505050565b600160a060020a033316600090815260016020526040812055565b600160a060020a0316600090815260016020526040902054151590565b61132b612256565b60008080808080871161133d57600080fd5b611346876121ab565b945084601e146113565784611359565b601d5b94508460010193505b601f8410156114a757600060048581548110151561137c57fe5b600091825260209091200154111561149c576113a46000600160048781548110151561098757fe5b92506004848154811015156113b557fe5b9060005260206000209001838154811015156113cd57fe5b90600052602060002090015491506003600083600019166000191681526020019081526020016000209050806001018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156114905780601f1061146557610100808354040283529160200191611490565b820191906000526020600020905b81548152906001019060200180831161147357829003601f168201915b505050505095506114b9565b600190930192611362565b60206040519081016040526000815295505b5050505050919050565b60036000826040518082805190602001908083835b602083106114f75780518252601f1990920191602091820191016114d8565b6001836020036101000a03801982511681845116179092525050509190910192506040915050519081900390208152602081019190915260400160002060090154151561154357600080fd5b600680546001810161155583826121ee565b91600052602060002090016000836040518082805190602001908083835b602083106115925780518252601f199092019160209182019101611573565b6001836020036101000a0380198251168184511617909252505050919091019250604091505051908190039020909155505050565b600160a060020a033316600090815260016020526040908190208054830190557fcf34ef537ac33ee1ac626ca1587a0a7e8e51561e5514f8cb36afa1c5102b3bab90516020808252600f908201527f50617920697320737563636573732100000000000000000000000000000000006040808301919091526060909101905180910390a150565b6116588282611b93565b5050565b50565b611667612256565b6003600060066116806000600160068054905003611ed9565b8154811061168a57fe5b906000526020600020900154600019166000191681526020019081526020016000206001018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156117435780601f1061171857610100808354040283529160200191611743565b820191906000526020600020905b81548152906001019060200180831161172657829003601f168201915b505050505090505b90565b60008060036000856040518082805190602001908083835b602083106117855780518252601f199092019160209182019101611766565b6001836020036101000a03801982511681845116179092525050509190910192506040915050519081900390208152602081019190915260400160009081206009810154909350116117d657600080fd5b82151561189357600582015415156118775760078054600684015542600484015580546001810161180783826121ee565b91600052602060002090016000866040518082805190602001908083835b602083106118445780518252601f199092019160209182019101611825565b6001836020036101000a038019825116818451161790925250505091909101925060409150505190819003902090915550505b610e108260040154420310156118935760058201805460010190555b610e108260040154420311806118ad575060018260050154115b15611b10576007546001901180156118d15750600754600683015460001990910114155b15611914576007805460001981019081106118e857fe5b9060005260206000209001546007836006015481548110151561190757fe5b6000918252602090912001555b60078054600019019061192790826121ee565b50600182600501541115611b015760078201805461ff0019166101001790556005546008830155600282015460048054909190811061196257fe5b9060005260206000209001826003015481548110151561197e57fe5b6000918252602082200155600282015460048054909190811061199d57fe5b60009182526020909120015490506001811180156119c2575060018103826003015414155b15611a3d57600482600201548154811015156119da57fe5b9060005260206000209001600182038154811015156119f557fe5b90600052602060002090015460048360020154815481101515611a1457fe5b90600052602060002090018360030154815481101515611a3057fe5b6000918252602090912001555b600160048360020154815481101515611a5257fe5b600091825260209091200180549190910390611a6e90826121ee565b506005805460018101611a8183826121ee565b91600052602060002090016000846001016040518082805460018160011615610100020316600290048015611aed5780601f10611acb576101008083540402835291820191611aed565b820191906000526020600020905b815481529060010190602001808311611ad9575b505091505060405190819003902090915550505b60006005830181905560048301555b50505050565b600060036000836040518082805190602001908083835b60208310611b4c5780518252601f199092019160209182019101611b2d565b6001836020036101000a038019825116818451161790925250505091909101925060409150505190819003902081526020810191909152604001600020600a015492915050565b600080611b9e612268565b60008411611bab57600080fd5b60036000866040518082805190602001908083835b60208310611bdf5780518252601f199092019160209182019101611bc0565b6001836020036101000a0380198251168184511617909252505050919091019250604091505051908190039020815260208101919091526040016000206009015415611c2a57600080fd5b600160a060020a03331660009081526002602052604081206001810154915481549195500203841115611c5c57600080fd5b60018301805485019055611c6f846121ab565b91506101806040519081016040528033600160a060020a03168152602001868152602001838152602001600484815481101515611ca857fe5b9060005260206000209001805490508152602001600081526020016000815260200160008152602001600015158152602001600015158152602001600081526020018581526020018581525090508060036000876040518082805190602001908083835b60208310611d2b5780518252601f199092019160209182019101611d0c565b6001836020036101000a0380198251168184511617909252505050919091019250604091505051908190039020815260208101919091526040016000208151815473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0391909116178155602082015181600101908051611dac9291602001906122d4565b5060408201518160020155606082015181600301556080820151816004015560a0820151816005015560c0820151816006015560e082015160078201805460ff19169115159190911790556101008201516007820180549115156101000261ff001990921691909117905561012082015181600801556101408201518160090155610160820151600a90910155506004805483908110611e4857fe5b6000918252602090912001805460018101611e6383826121ee565b91600052602060002090016000876040518082805190602001908083835b60208310611ea05780518252601f199092019160209182019101611e81565b6001836020036101000a038019825116818451161790925250505091909101925060409150505190819003902090915550505050505050565b60008282036001014233604051918252600160a060020a03166c01000000000000000000000000026020820152603401604051908190039020811515611f1b57fe5b068301905092915050565b600080600060036000856040518082805190602001908083835b60208310611f5f5780518252601f199092019160209182019101611f40565b6001836020036101000a038019825116818451161790925250505091909101925060409150505190819003902081526020810191909152604001600020600a810154909350611fad906121ab565b60028401549092508214611b105760048360020154815481101515611fce57fe5b90600052602060002090018360030154815481101515611fea57fe5b6000918252602082200155600283015460048054909190811061200957fe5b600091825260209091200154905060018111801561202e575060018103836003015414155b156120a9576004836002015481548110151561204657fe5b90600052602060002090016001820381548110151561206157fe5b9060005260206000209001546004846002015481548110151561208057fe5b9060005260206000209001846003015481548110151561209c57fe5b6000918252602090912001555b60018103600484600201548154811015156120c057fe5b9060005260206000209001816120d691906121ee565b506002830182905560048054839081106120ec57fe5b6000918252602090912001546003840155600480548390811061210b57fe5b600091825260209091200180546001810161212683826121ee565b916000526020600020900160008560010160405180828054600181600116156101000203166002900480156121925780601f10612170576101008083540402835291820191612192565b820191906000526020600020905b81548152906001019060200180831161217e575b5050915050604051908190039020909155505050505050565b60008080638000000084106121c357601e92506121e7565b506000905060015b8381116121e0576001909101906002026121cb565b6001820392505b5050919050565b81548183558181151161103257600083815260209020611032918101908301612352565b50805460018160011615610100020316600290046000825580601f10612238575061165c565b601f01602090049060005260206000209081019061165c9190612352565b60206040519081016040526000815290565b6101806040519081016040526000815260208101612284612256565b815260200160008152602001600081526020016000815260200160008152602001600081526020016000151581526020016000151581526020016000815260200160008152602001600081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061231557805160ff1916838001178555612342565b82800160010185558215612342579182015b82811115612342578251825591602001919060010190612327565b5061234e929150612352565b5090565b61174b91905b8082111561234e57600081556001016123585600a165627a7a7230582033f5a63bdfd452430dba0071a1edc77b9daa8eb55115dfdc0a9101a5ebe659a60029`

// DeployCasper deploys a new Ethereum contract, binding an instance of Casper to it.
func DeployCasper(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Casper, error) {
	parsed, err := abi.JSON(strings.NewReader(CasperABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CasperBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Casper{CasperCaller: CasperCaller{contract: contract}, CasperTransactor: CasperTransactor{contract: contract}, CasperFilterer: CasperFilterer{contract: contract}}, nil
}

// Casper is an auto generated Go binding around an Ethereum contract.
type Casper struct {
	CasperCaller     // Read-only binding to the contract
	CasperTransactor // Write-only binding to the contract
	CasperFilterer   // Log filterer for contract events
}

// CasperCaller is an auto generated read-only Go binding around an Ethereum contract.
type CasperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CasperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CasperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CasperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CasperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CasperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CasperSession struct {
	Contract     *Casper           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CasperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CasperCallerSession struct {
	Contract *CasperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// CasperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CasperTransactorSession struct {
	Contract     *CasperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CasperRaw is an auto generated low-level Go binding around an Ethereum contract.
type CasperRaw struct {
	Contract *Casper // Generic contract binding to access the raw methods on
}

// CasperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CasperCallerRaw struct {
	Contract *CasperCaller // Generic read-only contract binding to access the raw methods on
}

// CasperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CasperTransactorRaw struct {
	Contract *CasperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCasper creates a new instance of Casper, bound to a specific deployed contract.
func NewCasper(address common.Address, backend bind.ContractBackend) (*Casper, error) {
	contract, err := bindCasper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Casper{CasperCaller: CasperCaller{contract: contract}, CasperTransactor: CasperTransactor{contract: contract}, CasperFilterer: CasperFilterer{contract: contract}}, nil
}

// NewCasperCaller creates a new read-only instance of Casper, bound to a specific deployed contract.
func NewCasperCaller(address common.Address, caller bind.ContractCaller) (*CasperCaller, error) {
	contract, err := bindCasper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CasperCaller{contract: contract}, nil
}

// NewCasperTransactor creates a new write-only instance of Casper, bound to a specific deployed contract.
func NewCasperTransactor(address common.Address, transactor bind.ContractTransactor) (*CasperTransactor, error) {
	contract, err := bindCasper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CasperTransactor{contract: contract}, nil
}

// NewCasperFilterer creates a new log filterer instance of Casper, bound to a specific deployed contract.
func NewCasperFilterer(address common.Address, filterer bind.ContractFilterer) (*CasperFilterer, error) {
	contract, err := bindCasper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CasperFilterer{contract: contract}, nil
}

// bindCasper binds a generic wrapper to an already deployed contract.
func bindCasper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CasperABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Casper *CasperRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Casper.Contract.CasperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Casper *CasperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Casper.Contract.CasperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Casper *CasperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Casper.Contract.CasperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Casper *CasperCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Casper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Casper *CasperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Casper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Casper *CasperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Casper.Contract.contract.Transact(opts, method, params...)
}

// GetBootstrap is a free data retrieval call binding the contract method 0xd9290af9.
//
// Solidity: function getBootstrap() constant returns(ip string)
func (_Casper *CasperCaller) GetBootstrap(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Casper.contract.Call(opts, out, "getBootstrap")
	return *ret0, err
}

// GetBootstrap is a free data retrieval call binding the contract method 0xd9290af9.
//
// Solidity: function getBootstrap() constant returns(ip string)
func (_Casper *CasperSession) GetBootstrap() (string, error) {
	return _Casper.Contract.GetBootstrap(&_Casper.CallOpts)
}

// GetBootstrap is a free data retrieval call binding the contract method 0xd9290af9.
//
// Solidity: function getBootstrap() constant returns(ip string)
func (_Casper *CasperCallerSession) GetBootstrap() (string, error) {
	return _Casper.Contract.GetBootstrap(&_Casper.CallOpts)
}

// GetPeers is a free data retrieval call binding the contract method 0x9b92f9b1.
//
// Solidity: function getPeers(size uint256) constant returns(ip string)
func (_Casper *CasperCaller) GetPeers(opts *bind.CallOpts, size *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Casper.contract.Call(opts, out, "getPeers", size)
	return *ret0, err
}

// GetPeers is a free data retrieval call binding the contract method 0x9b92f9b1.
//
// Solidity: function getPeers(size uint256) constant returns(ip string)
func (_Casper *CasperSession) GetPeers(size *big.Int) (string, error) {
	return _Casper.Contract.GetPeers(&_Casper.CallOpts, size)
}

// GetPeers is a free data retrieval call binding the contract method 0x9b92f9b1.
//
// Solidity: function getPeers(size uint256) constant returns(ip string)
func (_Casper *CasperCallerSession) GetPeers(size *big.Int) (string, error) {
	return _Casper.Contract.GetPeers(&_Casper.CallOpts, size)
}

// IsPrepaid is a free data retrieval call binding the contract method 0x73e07e14.
//
// Solidity: function isPrepaid(client address) constant returns(prepaid bool)
func (_Casper *CasperCaller) IsPrepaid(opts *bind.CallOpts, client common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Casper.contract.Call(opts, out, "isPrepaid", client)
	return *ret0, err
}

// IsPrepaid is a free data retrieval call binding the contract method 0x73e07e14.
//
// Solidity: function isPrepaid(client address) constant returns(prepaid bool)
func (_Casper *CasperSession) IsPrepaid(client common.Address) (bool, error) {
	return _Casper.Contract.IsPrepaid(&_Casper.CallOpts, client)
}

// IsPrepaid is a free data retrieval call binding the contract method 0x73e07e14.
//
// Solidity: function isPrepaid(client address) constant returns(prepaid bool)
func (_Casper *CasperCallerSession) IsPrepaid(client common.Address) (bool, error) {
	return _Casper.Contract.IsPrepaid(&_Casper.CallOpts, client)
}

// MachineInformation is a free data retrieval call binding the contract method 0xf8c5df33.
//
// Solidity: function machineInformation(ip string) constant returns(uint256)
func (_Casper *CasperCaller) MachineInformation(opts *bind.CallOpts, ip string) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Casper.contract.Call(opts, out, "machineInformation", ip)
	return *ret0, err
}

// MachineInformation is a free data retrieval call binding the contract method 0xf8c5df33.
//
// Solidity: function machineInformation(ip string) constant returns(uint256)
func (_Casper *CasperSession) MachineInformation(ip string) (*big.Int, error) {
	return _Casper.Contract.MachineInformation(&_Casper.CallOpts, ip)
}

// MachineInformation is a free data retrieval call binding the contract method 0xf8c5df33.
//
// Solidity: function machineInformation(ip string) constant returns(uint256)
func (_Casper *CasperCallerSession) MachineInformation(ip string) (*big.Int, error) {
	return _Casper.Contract.MachineInformation(&_Casper.CallOpts, ip)
}

// AddProviderMachine is a paid mutator transaction binding the contract method 0xfde18f47.
//
// Solidity: function addProviderMachine(ip string, size uint256) returns()
func (_Casper *CasperTransactor) AddProviderMachine(opts *bind.TransactOpts, ip string, size *big.Int) (*types.Transaction, error) {
	return _Casper.contract.Transact(opts, "addProviderMachine", ip, size)
}

// AddProviderMachine is a paid mutator transaction binding the contract method 0xfde18f47.
//
// Solidity: function addProviderMachine(ip string, size uint256) returns()
func (_Casper *CasperSession) AddProviderMachine(ip string, size *big.Int) (*types.Transaction, error) {
	return _Casper.Contract.AddProviderMachine(&_Casper.TransactOpts, ip, size)
}

// AddProviderMachine is a paid mutator transaction binding the contract method 0xfde18f47.
//
// Solidity: function addProviderMachine(ip string, size uint256) returns()
func (_Casper *CasperTransactorSession) AddProviderMachine(ip string, size *big.Int) (*types.Transaction, error) {
	return _Casper.Contract.AddProviderMachine(&_Casper.TransactOpts, ip, size)
}

// AddToken is a paid mutator transaction binding the contract method 0x179d375c.
//
// Solidity: function addToken(amount uint256) returns()
func (_Casper *CasperTransactor) AddToken(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Casper.contract.Transact(opts, "addToken", amount)
}

// AddToken is a paid mutator transaction binding the contract method 0x179d375c.
//
// Solidity: function addToken(amount uint256) returns()
func (_Casper *CasperSession) AddToken(amount *big.Int) (*types.Transaction, error) {
	return _Casper.Contract.AddToken(&_Casper.TransactOpts, amount)
}

// AddToken is a paid mutator transaction binding the contract method 0x179d375c.
//
// Solidity: function addToken(amount uint256) returns()
func (_Casper *CasperTransactorSession) AddToken(amount *big.Int) (*types.Transaction, error) {
	return _Casper.Contract.AddToken(&_Casper.TransactOpts, amount)
}

// ClearPrepay is a paid mutator transaction binding the contract method 0x6e2e1887.
//
// Solidity: function clearPrepay() returns()
func (_Casper *CasperTransactor) ClearPrepay(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Casper.contract.Transact(opts, "clearPrepay")
}

// ClearPrepay is a paid mutator transaction binding the contract method 0x6e2e1887.
//
// Solidity: function clearPrepay() returns()
func (_Casper *CasperSession) ClearPrepay() (*types.Transaction, error) {
	return _Casper.Contract.ClearPrepay(&_Casper.TransactOpts)
}

// ClearPrepay is a paid mutator transaction binding the contract method 0x6e2e1887.
//
// Solidity: function clearPrepay() returns()
func (_Casper *CasperTransactorSession) ClearPrepay() (*types.Transaction, error) {
	return _Casper.Contract.ClearPrepay(&_Casper.TransactOpts)
}

// ConfirmDownload is a paid mutator transaction binding the contract method 0x28748e91.
//
// Solidity: function confirmDownload() returns()
func (_Casper *CasperTransactor) ConfirmDownload(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Casper.contract.Transact(opts, "confirmDownload")
}

// ConfirmDownload is a paid mutator transaction binding the contract method 0x28748e91.
//
// Solidity: function confirmDownload() returns()
func (_Casper *CasperSession) ConfirmDownload() (*types.Transaction, error) {
	return _Casper.Contract.ConfirmDownload(&_Casper.TransactOpts)
}

// ConfirmDownload is a paid mutator transaction binding the contract method 0x28748e91.
//
// Solidity: function confirmDownload() returns()
func (_Casper *CasperTransactorSession) ConfirmDownload() (*types.Transaction, error) {
	return _Casper.Contract.ConfirmDownload(&_Casper.TransactOpts)
}

// ConfirmTransfer is a paid mutator transaction binding the contract method 0xbaa49301.
//
// Solidity: function confirmTransfer() returns()
func (_Casper *CasperTransactor) ConfirmTransfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Casper.contract.Transact(opts, "confirmTransfer")
}

// ConfirmTransfer is a paid mutator transaction binding the contract method 0xbaa49301.
//
// Solidity: function confirmTransfer() returns()
func (_Casper *CasperSession) ConfirmTransfer() (*types.Transaction, error) {
	return _Casper.Contract.ConfirmTransfer(&_Casper.TransactOpts)
}

// ConfirmTransfer is a paid mutator transaction binding the contract method 0xbaa49301.
//
// Solidity: function confirmTransfer() returns()
func (_Casper *CasperTransactorSession) ConfirmTransfer() (*types.Transaction, error) {
	return _Casper.Contract.ConfirmTransfer(&_Casper.TransactOpts)
}

// ConfirmUpload is a paid mutator transaction binding the contract method 0x3f65a6d0.
//
// Solidity: function confirmUpload(ip string, size uint256) returns()
func (_Casper *CasperTransactor) ConfirmUpload(opts *bind.TransactOpts, ip string, size *big.Int) (*types.Transaction, error) {
	return _Casper.contract.Transact(opts, "confirmUpload", ip, size)
}

// ConfirmUpload is a paid mutator transaction binding the contract method 0x3f65a6d0.
//
// Solidity: function confirmUpload(ip string, size uint256) returns()
func (_Casper *CasperSession) ConfirmUpload(ip string, size *big.Int) (*types.Transaction, error) {
	return _Casper.Contract.ConfirmUpload(&_Casper.TransactOpts, ip, size)
}

// ConfirmUpload is a paid mutator transaction binding the contract method 0x3f65a6d0.
//
// Solidity: function confirmUpload(ip string, size uint256) returns()
func (_Casper *CasperTransactorSession) ConfirmUpload(ip string, size *big.Int) (*types.Transaction, error) {
	return _Casper.Contract.ConfirmUpload(&_Casper.TransactOpts, ip, size)
}

// GetPingTarget is a paid mutator transaction binding the contract method 0x1acaa833.
//
// Solidity: function getPingTarget(srcIp string) returns()
func (_Casper *CasperTransactor) GetPingTarget(opts *bind.TransactOpts, srcIp string) (*types.Transaction, error) {
	return _Casper.contract.Transact(opts, "getPingTarget", srcIp)
}

// GetPingTarget is a paid mutator transaction binding the contract method 0x1acaa833.
//
// Solidity: function getPingTarget(srcIp string) returns()
func (_Casper *CasperSession) GetPingTarget(srcIp string) (*types.Transaction, error) {
	return _Casper.Contract.GetPingTarget(&_Casper.TransactOpts, srcIp)
}

// GetPingTarget is a paid mutator transaction binding the contract method 0x1acaa833.
//
// Solidity: function getPingTarget(srcIp string) returns()
func (_Casper *CasperTransactorSession) GetPingTarget(srcIp string) (*types.Transaction, error) {
	return _Casper.Contract.GetPingTarget(&_Casper.TransactOpts, srcIp)
}

// NotifyDelete is a paid mutator transaction binding the contract method 0xc72f7c54.
//
// Solidity: function notifyDelete(uuid uint128) returns()
func (_Casper *CasperTransactor) NotifyDelete(opts *bind.TransactOpts, uuid *big.Int) (*types.Transaction, error) {
	return _Casper.contract.Transact(opts, "notifyDelete", uuid)
}

// NotifyDelete is a paid mutator transaction binding the contract method 0xc72f7c54.
//
// Solidity: function notifyDelete(uuid uint128) returns()
func (_Casper *CasperSession) NotifyDelete(uuid *big.Int) (*types.Transaction, error) {
	return _Casper.Contract.NotifyDelete(&_Casper.TransactOpts, uuid)
}

// NotifyDelete is a paid mutator transaction binding the contract method 0xc72f7c54.
//
// Solidity: function notifyDelete(uuid uint128) returns()
func (_Casper *CasperTransactorSession) NotifyDelete(uuid *big.Int) (*types.Transaction, error) {
	return _Casper.Contract.NotifyDelete(&_Casper.TransactOpts, uuid)
}

// NotifySpaceFreed is a paid mutator transaction binding the contract method 0x4d2ab7e4.
//
// Solidity: function notifySpaceFreed(ip string, size uint256) returns()
func (_Casper *CasperTransactor) NotifySpaceFreed(opts *bind.TransactOpts, ip string, size *big.Int) (*types.Transaction, error) {
	return _Casper.contract.Transact(opts, "notifySpaceFreed", ip, size)
}

// NotifySpaceFreed is a paid mutator transaction binding the contract method 0x4d2ab7e4.
//
// Solidity: function notifySpaceFreed(ip string, size uint256) returns()
func (_Casper *CasperSession) NotifySpaceFreed(ip string, size *big.Int) (*types.Transaction, error) {
	return _Casper.Contract.NotifySpaceFreed(&_Casper.TransactOpts, ip, size)
}

// NotifySpaceFreed is a paid mutator transaction binding the contract method 0x4d2ab7e4.
//
// Solidity: function notifySpaceFreed(ip string, size uint256) returns()
func (_Casper *CasperTransactorSession) NotifySpaceFreed(ip string, size *big.Int) (*types.Transaction, error) {
	return _Casper.Contract.NotifySpaceFreed(&_Casper.TransactOpts, ip, size)
}

// PrePay is a paid mutator transaction binding the contract method 0xaac80f47.
//
// Solidity: function prePay(pay uint256) returns()
func (_Casper *CasperTransactor) PrePay(opts *bind.TransactOpts, pay *big.Int) (*types.Transaction, error) {
	return _Casper.contract.Transact(opts, "prePay", pay)
}

// PrePay is a paid mutator transaction binding the contract method 0xaac80f47.
//
// Solidity: function prePay(pay uint256) returns()
func (_Casper *CasperSession) PrePay(pay *big.Int) (*types.Transaction, error) {
	return _Casper.Contract.PrePay(&_Casper.TransactOpts, pay)
}

// PrePay is a paid mutator transaction binding the contract method 0xaac80f47.
//
// Solidity: function prePay(pay uint256) returns()
func (_Casper *CasperTransactorSession) PrePay(pay *big.Int) (*types.Transaction, error) {
	return _Casper.Contract.PrePay(&_Casper.TransactOpts, pay)
}

// RegisterProvider is a paid mutator transaction binding the contract method 0xbb66e09d.
//
// Solidity: function registerProvider(ip string, size uint256) returns()
func (_Casper *CasperTransactor) RegisterProvider(opts *bind.TransactOpts, ip string, size *big.Int) (*types.Transaction, error) {
	return _Casper.contract.Transact(opts, "registerProvider", ip, size)
}

// RegisterProvider is a paid mutator transaction binding the contract method 0xbb66e09d.
//
// Solidity: function registerProvider(ip string, size uint256) returns()
func (_Casper *CasperSession) RegisterProvider(ip string, size *big.Int) (*types.Transaction, error) {
	return _Casper.Contract.RegisterProvider(&_Casper.TransactOpts, ip, size)
}

// RegisterProvider is a paid mutator transaction binding the contract method 0xbb66e09d.
//
// Solidity: function registerProvider(ip string, size uint256) returns()
func (_Casper *CasperTransactorSession) RegisterProvider(ip string, size *big.Int) (*types.Transaction, error) {
	return _Casper.Contract.RegisterProvider(&_Casper.TransactOpts, ip, size)
}

// RemoveProviderMachine is a paid mutator transaction binding the contract method 0x5107e4ee.
//
// Solidity: function removeProviderMachine(ip string) returns()
func (_Casper *CasperTransactor) RemoveProviderMachine(opts *bind.TransactOpts, ip string) (*types.Transaction, error) {
	return _Casper.contract.Transact(opts, "removeProviderMachine", ip)
}

// RemoveProviderMachine is a paid mutator transaction binding the contract method 0x5107e4ee.
//
// Solidity: function removeProviderMachine(ip string) returns()
func (_Casper *CasperSession) RemoveProviderMachine(ip string) (*types.Transaction, error) {
	return _Casper.Contract.RemoveProviderMachine(&_Casper.TransactOpts, ip)
}

// RemoveProviderMachine is a paid mutator transaction binding the contract method 0x5107e4ee.
//
// Solidity: function removeProviderMachine(ip string) returns()
func (_Casper *CasperTransactorSession) RemoveProviderMachine(ip string) (*types.Transaction, error) {
	return _Casper.Contract.RemoveProviderMachine(&_Casper.TransactOpts, ip)
}

// SendPingResult is a paid mutator transaction binding the contract method 0xdc3dd578.
//
// Solidity: function sendPingResult(ip string, success bool) returns()
func (_Casper *CasperTransactor) SendPingResult(opts *bind.TransactOpts, ip string, success bool) (*types.Transaction, error) {
	return _Casper.contract.Transact(opts, "sendPingResult", ip, success)
}

// SendPingResult is a paid mutator transaction binding the contract method 0xdc3dd578.
//
// Solidity: function sendPingResult(ip string, success bool) returns()
func (_Casper *CasperSession) SendPingResult(ip string, success bool) (*types.Transaction, error) {
	return _Casper.Contract.SendPingResult(&_Casper.TransactOpts, ip, success)
}

// SendPingResult is a paid mutator transaction binding the contract method 0xdc3dd578.
//
// Solidity: function sendPingResult(ip string, success bool) returns()
func (_Casper *CasperTransactorSession) SendPingResult(ip string, success bool) (*types.Transaction, error) {
	return _Casper.Contract.SendPingResult(&_Casper.TransactOpts, ip, success)
}

// SetBootstrap is a paid mutator transaction binding the contract method 0xa746dadb.
//
// Solidity: function setBootstrap(ip string) returns()
func (_Casper *CasperTransactor) SetBootstrap(opts *bind.TransactOpts, ip string) (*types.Transaction, error) {
	return _Casper.contract.Transact(opts, "setBootstrap", ip)
}

// SetBootstrap is a paid mutator transaction binding the contract method 0xa746dadb.
//
// Solidity: function setBootstrap(ip string) returns()
func (_Casper *CasperSession) SetBootstrap(ip string) (*types.Transaction, error) {
	return _Casper.Contract.SetBootstrap(&_Casper.TransactOpts, ip)
}

// SetBootstrap is a paid mutator transaction binding the contract method 0xa746dadb.
//
// Solidity: function setBootstrap(ip string) returns()
func (_Casper *CasperTransactorSession) SetBootstrap(ip string) (*types.Transaction, error) {
	return _Casper.Contract.SetBootstrap(&_Casper.TransactOpts, ip)
}

// CasperLogIterator is returned from FilterLog and is used to iterate over the raw logs and unpacked data for Log events raised by the Casper contract.
type CasperLogIterator struct {
	Event *CasperLog // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CasperLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CasperLog)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CasperLog)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CasperLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CasperLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CasperLog represents a Log event raised by the Casper contract.
type CasperLog struct {
	Info string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLog is a free log retrieval operation binding the contract event 0xcf34ef537ac33ee1ac626ca1587a0a7e8e51561e5514f8cb36afa1c5102b3bab.
//
// Solidity: event Log(info string)
func (_Casper *CasperFilterer) FilterLog(opts *bind.FilterOpts) (*CasperLogIterator, error) {

	logs, sub, err := _Casper.contract.FilterLogs(opts, "Log")
	if err != nil {
		return nil, err
	}
	return &CasperLogIterator{contract: _Casper.contract, event: "Log", logs: logs, sub: sub}, nil
}

// WatchLog is a free log subscription operation binding the contract event 0xcf34ef537ac33ee1ac626ca1587a0a7e8e51561e5514f8cb36afa1c5102b3bab.
//
// Solidity: event Log(info string)
func (_Casper *CasperFilterer) WatchLog(opts *bind.WatchOpts, sink chan<- *CasperLog) (event.Subscription, error) {

	logs, sub, err := _Casper.contract.WatchLogs(opts, "Log")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CasperLog)
				if err := _Casper.contract.UnpackLog(event, "Log", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// CasperProviderCheckEventIterator is returned from FilterProviderCheckEvent and is used to iterate over the raw logs and unpacked data for ProviderCheckEvent events raised by the Casper contract.
type CasperProviderCheckEventIterator struct {
	Event *CasperProviderCheckEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CasperProviderCheckEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CasperProviderCheckEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CasperProviderCheckEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CasperProviderCheckEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CasperProviderCheckEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CasperProviderCheckEvent represents a ProviderCheckEvent event raised by the Casper contract.
type CasperProviderCheckEvent struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterProviderCheckEvent is a free log retrieval operation binding the contract event 0x3ba3d58761f64d3f96a3d0808397c076d4936a34f18b36fb0b297a93ee8a83b9.
//
// Solidity: event ProviderCheckEvent()
func (_Casper *CasperFilterer) FilterProviderCheckEvent(opts *bind.FilterOpts) (*CasperProviderCheckEventIterator, error) {

	logs, sub, err := _Casper.contract.FilterLogs(opts, "ProviderCheckEvent")
	if err != nil {
		return nil, err
	}
	return &CasperProviderCheckEventIterator{contract: _Casper.contract, event: "ProviderCheckEvent", logs: logs, sub: sub}, nil
}

// WatchProviderCheckEvent is a free log subscription operation binding the contract event 0x3ba3d58761f64d3f96a3d0808397c076d4936a34f18b36fb0b297a93ee8a83b9.
//
// Solidity: event ProviderCheckEvent()
func (_Casper *CasperFilterer) WatchProviderCheckEvent(opts *bind.WatchOpts, sink chan<- *CasperProviderCheckEvent) (event.Subscription, error) {

	logs, sub, err := _Casper.contract.WatchLogs(opts, "ProviderCheckEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CasperProviderCheckEvent)
				if err := _Casper.contract.UnpackLog(event, "ProviderCheckEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// CasperReturnStringIterator is returned from FilterReturnString and is used to iterate over the raw logs and unpacked data for ReturnString events raised by the Casper contract.
type CasperReturnStringIterator struct {
	Event *CasperReturnString // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CasperReturnStringIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CasperReturnString)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CasperReturnString)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CasperReturnStringIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CasperReturnStringIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CasperReturnString represents a ReturnString event raised by the Casper contract.
type CasperReturnString struct {
	Val string
	Raw types.Log // Blockchain specific contextual infos
}

// FilterReturnString is a free log retrieval operation binding the contract event 0x9e91c0d67653515860c12ff5c5211b3ef5b98565f70728283708dd7bf754c4f1.
//
// Solidity: event ReturnString(val string)
func (_Casper *CasperFilterer) FilterReturnString(opts *bind.FilterOpts) (*CasperReturnStringIterator, error) {

	logs, sub, err := _Casper.contract.FilterLogs(opts, "ReturnString")
	if err != nil {
		return nil, err
	}
	return &CasperReturnStringIterator{contract: _Casper.contract, event: "ReturnString", logs: logs, sub: sub}, nil
}

// WatchReturnString is a free log subscription operation binding the contract event 0x9e91c0d67653515860c12ff5c5211b3ef5b98565f70728283708dd7bf754c4f1.
//
// Solidity: event ReturnString(val string)
func (_Casper *CasperFilterer) WatchReturnString(opts *bind.WatchOpts, sink chan<- *CasperReturnString) (event.Subscription, error) {

	logs, sub, err := _Casper.contract.WatchLogs(opts, "ReturnString")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CasperReturnString)
				if err := _Casper.contract.UnpackLog(event, "ReturnString", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

