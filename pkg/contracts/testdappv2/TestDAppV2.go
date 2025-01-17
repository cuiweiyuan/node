// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package testdappv2

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// TestDAppV2MessageContext is an auto generated low-level Go binding around an user-defined struct.
type TestDAppV2MessageContext struct {
	Sender common.Address
}

// TestDAppV2RevertContext is an auto generated low-level Go binding around an user-defined struct.
type TestDAppV2RevertContext struct {
	Sender        common.Address
	Asset         common.Address
	Amount        *big.Int
	RevertMessage []byte
}

// TestDAppV2zContext is an auto generated low-level Go binding around an user-defined struct.
type TestDAppV2zContext struct {
	Origin  []byte
	Sender  common.Address
	ChainID *big.Int
}

// TestDAppV2MetaData contains all meta data concerning the TestDAppV2 contract.
var TestDAppV2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"isZetaChain_\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"gateway_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"NO_MESSAGE_CALL\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"amountWithMessage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"calledWithMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"erc20\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"erc20Call\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"gasCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gateway\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"gatewayCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"}],\"name\":\"gatewayDeposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"gatewayDepositAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"getAmountWithMessage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"getCalledWithMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"getNoMessageIndex\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isZetaChain\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"origin\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"internalType\":\"structTestDAppV2.zContext\",\"name\":\"_context\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_zrc20\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"onCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"internalType\":\"structTestDAppV2.MessageContext\",\"name\":\"messageContext\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"onCall\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"revertMessage\",\"type\":\"bytes\"}],\"internalType\":\"structTestDAppV2.RevertContext\",\"name\":\"revertContext\",\"type\":\"tuple\"}],\"name\":\"onRevert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"senderWithMessage\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"simpleCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60c060405234801561001057600080fd5b50604051611ebb380380611ebb83398181016040528101906100329190610114565b8115156080811515815250508073ffffffffffffffffffffffffffffffffffffffff1660a08173ffffffffffffffffffffffffffffffffffffffff16815250505050610154565b600080fd5b60008115159050919050565b6100938161007e565b811461009e57600080fd5b50565b6000815190506100b08161008a565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006100e1826100b6565b9050919050565b6100f1816100d6565b81146100fc57600080fd5b50565b60008151905061010e816100e8565b92915050565b6000806040838503121561012b5761012a610079565b5b6000610139858286016100a1565b925050602061014a858286016100ff565b9150509250929050565b60805160a051611d116101aa6000396000818161045b015281816104d40152818161084e0152610c730152600081816104a90152818161082301528181610abb01528181610c240152610c480152611d116000f3fe60806040526004361061010d5760003560e01c8063ad23b28b11610095578063c91f356711610064578063c91f35671461035b578063deb3b1e414610386578063e2842ed7146103a2578063f592cbfb146103df578063f936ae851461041c57610114565b8063ad23b28b146102a1578063c7a339a9146102de578063c85f843414610307578063c9028a361461033257610114565b80635bcfd616116100dc5780635bcfd616146101d3578063676cc054146101fc5780639291fe261461022c5780639ca016ed14610269578063a799911f1461028557610114565b8063116191b61461011957806336e980a01461014457806341a3cd4a1461016d5780634297a2631461019657610114565b3661011457005b600080fd5b34801561012557600080fd5b5061012e610459565b60405161013b9190611005565b60405180910390f35b34801561015057600080fd5b5061016b6004803603810190610166919061117a565b61047d565b005b34801561017957600080fd5b50610194600480360381019061018f919061124f565b6104a7565b005b3480156101a257600080fd5b506101bd60048036038101906101b891906112e5565b6105ce565b6040516101ca919061132b565b60405180910390f35b3480156101df57600080fd5b506101fa60048036038101906101f59190611396565b6105e6565b005b61021660048036038101906102119190611459565b6106cc565b6040516102239190611538565b60405180910390f35b34801561023857600080fd5b50610253600480360381019061024e919061117a565b6107de565b604051610260919061132b565b60405180910390f35b610283600480360381019061027e919061155a565b610821565b005b61029f600480360381019061029a919061117a565b610943565b005b3480156102ad57600080fd5b506102c860048036038101906102c3919061155a565b61096c565b6040516102d591906115dc565b60405180910390f35b3480156102ea57600080fd5b506103056004803603810190610300919061163c565b6109cc565b005b34801561031357600080fd5b5061031c610a80565b60405161032991906115dc565b60405180910390f35b34801561033e57600080fd5b50610359600480360381019061035491906116ca565b610ab9565b005b34801561036757600080fd5b50610370610c22565b60405161037d919061172e565b60405180910390f35b6103a0600480360381019061039b919061124f565b610c46565b005b3480156103ae57600080fd5b506103c960048036038101906103c491906112e5565b610d6e565b6040516103d6919061172e565b60405180910390f35b3480156103eb57600080fd5b506104066004803603810190610401919061117a565b610d8e565b604051610413919061172e565b60405180910390f35b34801561042857600080fd5b50610443600480360381019061043e91906117ea565b610dde565b6040516104509190611005565b60405180910390f35b7f000000000000000000000000000000000000000000000000000000000000000081565b61048681610e27565b1561049057600080fd5b61049981610e7d565b6104a4816000610ed1565b50565b7f0000000000000000000000000000000000000000000000000000000000000000156104d257600080fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16631becceb48484846040518060a001604052803373ffffffffffffffffffffffffffffffffffffffff168152602001600015158152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200160405180602001604052806000815250815260200160008152506040518563ffffffff1660e01b8152600401610597949392919061194d565b600060405180830381600087803b1580156105b157600080fd5b505af11580156105c5573d6000803e3d6000fd5b50505050505050565b60036020528060005260406000206000915090505481565b61063382828080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050610e27565b1561063d57600080fd5b60008083839050146106935782828080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050506106af565b6106ae8660200160208101906106a9919061155a565b61096c565b5b90506106ba81610e7d565b6106c48185610ed1565b505050505050565b606060008084849050146107245783838080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050610740565b61073f85600001602081019061073a919061155a565b61096c565b5b905061074b81610e7d565b6107558134610ed1565b846000016020810190610768919061155a565b60028260405161077891906119d0565b908152602001604051809103902060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550604051806020016040528060008152509150509392505050565b600060036000836040516020016107f59190611a23565b604051602081830303815290604052805190602001208152602001908152602001600020549050919050565b7f00000000000000000000000000000000000000000000000000000000000000001561084c57600080fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663726ac97c34836040518060a001604052803373ffffffffffffffffffffffffffffffffffffffff168152602001600015158152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200160405180602001604052806000815250815260200160008152506040518463ffffffff1660e01b815260040161090e929190611a3a565b6000604051808303818588803b15801561092757600080fd5b505af115801561093b573d6000803e3d6000fd5b505050505050565b61094c81610e27565b1561095657600080fd5b61095f81610e7d565b6109698134610ed1565b50565b60606040518060400160405280601681526020017f63616c6c65642077697468206e6f206d65737361676500000000000000000000815250826040516020016109b6929190611ab2565b6040516020818303038152906040529050919050565b6109d581610e27565b156109df57600080fd5b8273ffffffffffffffffffffffffffffffffffffffff166323b872dd3330856040518463ffffffff1660e01b8152600401610a1c93929190611ada565b6020604051808303816000875af1158015610a3b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a5f9190611b3d565b610a6857600080fd5b610a7181610e7d565b610a7b8183610ed1565b505050565b6040518060400160405280601681526020017f63616c6c65642077697468206e6f206d6573736167650000000000000000000081525081565b7f000000000000000000000000000000000000000000000000000000000000000015610ae857610ae7610f13565b5b610b43818060600190610afb9190611b79565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050610e7d565b610ba0818060600190610b569190611b79565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050506000610ed1565b806000016020810190610bb3919061155a565b6002828060600190610bc59190611b79565b604051610bd3929190611c01565b908152602001604051809103902060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b7f000000000000000000000000000000000000000000000000000000000000000081565b7f000000000000000000000000000000000000000000000000000000000000000015610c7157600080fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663744b9b8b348585856040518060a001604052803373ffffffffffffffffffffffffffffffffffffffff168152602001600015158152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200160405180602001604052806000815250815260200160008152506040518663ffffffff1660e01b8152600401610d37949392919061194d565b6000604051808303818588803b158015610d5057600080fd5b505af1158015610d64573d6000803e3d6000fd5b5050505050505050565b60016020528060005260406000206000915054906101000a900460ff1681565b60006001600083604051602001610da59190611a23565b60405160208183030381529060405280519060200120815260200190815260200160002060009054906101000a900460ff169050919050565b6002818051602081018201805184825260208301602085012081835280955050505050506000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000604051602001610e3890611c66565b6040516020818303038152906040528051906020012082604051602001610e5f9190611a23565b60405160208183030381529060405280519060200120149050919050565b600180600083604051602001610e939190611a23565b60405160208183030381529060405280519060200120815260200190815260200160002060006101000a81548160ff02191690831515021790555050565b806003600084604051602001610ee79190611a23565b604051602081830303815290604052805190602001208152602001908152602001600020819055505050565b60006207a12090506000614e20905060008183610f309190611caa565b905060005b81811015610f735760008190806001815401808255809150506001900390600052602060002001600090919091909150558080600101915050610f35565b50600080610f819190610f86565b505050565b5080546000825590600052602060002090810190610fa49190610fa7565b50565b5b80821115610fc0576000816000905550600101610fa8565b5090565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610fef82610fc4565b9050919050565b610fff81610fe4565b82525050565b600060208201905061101a6000830184610ff6565b92915050565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6110878261103e565b810181811067ffffffffffffffff821117156110a6576110a561104f565b5b80604052505050565b60006110b9611020565b90506110c5828261107e565b919050565b600067ffffffffffffffff8211156110e5576110e461104f565b5b6110ee8261103e565b9050602081019050919050565b82818337600083830152505050565b600061111d611118846110ca565b6110af565b90508281526020810184848401111561113957611138611039565b5b6111448482856110fb565b509392505050565b600082601f83011261116157611160611034565b5b813561117184826020860161110a565b91505092915050565b6000602082840312156111905761118f61102a565b5b600082013567ffffffffffffffff8111156111ae576111ad61102f565b5b6111ba8482850161114c565b91505092915050565b6111cc81610fe4565b81146111d757600080fd5b50565b6000813590506111e9816111c3565b92915050565b600080fd5b600080fd5b60008083601f84011261120f5761120e611034565b5b8235905067ffffffffffffffff81111561122c5761122b6111ef565b5b602083019150836001820283011115611248576112476111f4565b5b9250929050565b6000806000604084860312156112685761126761102a565b5b6000611276868287016111da565b935050602084013567ffffffffffffffff8111156112975761129661102f565b5b6112a3868287016111f9565b92509250509250925092565b6000819050919050565b6112c2816112af565b81146112cd57600080fd5b50565b6000813590506112df816112b9565b92915050565b6000602082840312156112fb576112fa61102a565b5b6000611309848285016112d0565b91505092915050565b6000819050919050565b61132581611312565b82525050565b6000602082019050611340600083018461131c565b92915050565b600080fd5b60006060828403121561136157611360611346565b5b81905092915050565b61137381611312565b811461137e57600080fd5b50565b6000813590506113908161136a565b92915050565b6000806000806000608086880312156113b2576113b161102a565b5b600086013567ffffffffffffffff8111156113d0576113cf61102f565b5b6113dc8882890161134b565b95505060206113ed888289016111da565b94505060406113fe88828901611381565b935050606086013567ffffffffffffffff81111561141f5761141e61102f565b5b61142b888289016111f9565b92509250509295509295909350565b6000602082840312156114505761144f611346565b5b81905092915050565b6000806000604084860312156114725761147161102a565b5b60006114808682870161143a565b935050602084013567ffffffffffffffff8111156114a1576114a061102f565b5b6114ad868287016111f9565b92509250509250925092565b600081519050919050565b600082825260208201905092915050565b60005b838110156114f35780820151818401526020810190506114d8565b60008484015250505050565b600061150a826114b9565b61151481856114c4565b93506115248185602086016114d5565b61152d8161103e565b840191505092915050565b6000602082019050818103600083015261155281846114ff565b905092915050565b6000602082840312156115705761156f61102a565b5b600061157e848285016111da565b91505092915050565b600081519050919050565b600082825260208201905092915050565b60006115ae82611587565b6115b88185611592565b93506115c88185602086016114d5565b6115d18161103e565b840191505092915050565b600060208201905081810360008301526115f681846115a3565b905092915050565b600061160982610fe4565b9050919050565b611619816115fe565b811461162457600080fd5b50565b60008135905061163681611610565b92915050565b6000806000606084860312156116555761165461102a565b5b600061166386828701611627565b935050602061167486828701611381565b925050604084013567ffffffffffffffff8111156116955761169461102f565b5b6116a18682870161114c565b9150509250925092565b6000608082840312156116c1576116c0611346565b5b81905092915050565b6000602082840312156116e0576116df61102a565b5b600082013567ffffffffffffffff8111156116fe576116fd61102f565b5b61170a848285016116ab565b91505092915050565b60008115159050919050565b61172881611713565b82525050565b6000602082019050611743600083018461171f565b92915050565b600067ffffffffffffffff8211156117645761176361104f565b5b61176d8261103e565b9050602081019050919050565b600061178d61178884611749565b6110af565b9050828152602081018484840111156117a9576117a8611039565b5b6117b48482856110fb565b509392505050565b600082601f8301126117d1576117d0611034565b5b81356117e184826020860161177a565b91505092915050565b600060208284031215611800576117ff61102a565b5b600082013567ffffffffffffffff81111561181e5761181d61102f565b5b61182a848285016117bc565b91505092915050565b600061183f83856114c4565b935061184c8385846110fb565b6118558361103e565b840190509392505050565b61186981610fe4565b82525050565b61187881611713565b82525050565b600082825260208201905092915050565b600061189a826114b9565b6118a4818561187e565b93506118b48185602086016114d5565b6118bd8161103e565b840191505092915050565b6118d181611312565b82525050565b600060a0830160008301516118ef6000860182611860565b506020830151611902602086018261186f565b5060408301516119156040860182611860565b506060830151848203606086015261192d828261188f565b915050608083015161194260808601826118c8565b508091505092915050565b60006060820190506119626000830187610ff6565b8181036020830152611975818587611833565b9050818103604083015261198981846118d7565b905095945050505050565b600081905092915050565b60006119aa826114b9565b6119b48185611994565b93506119c48185602086016114d5565b80840191505092915050565b60006119dc828461199f565b915081905092915050565b600081905092915050565b60006119fd82611587565b611a0781856119e7565b9350611a178185602086016114d5565b80840191505092915050565b6000611a2f82846119f2565b915081905092915050565b6000604082019050611a4f6000830185610ff6565b8181036020830152611a6181846118d7565b90509392505050565b60008160601b9050919050565b6000611a8282611a6a565b9050919050565b6000611a9482611a77565b9050919050565b611aac611aa782610fe4565b611a89565b82525050565b6000611abe82856119f2565b9150611aca8284611a9b565b6014820191508190509392505050565b6000606082019050611aef6000830186610ff6565b611afc6020830185610ff6565b611b09604083018461131c565b949350505050565b611b1a81611713565b8114611b2557600080fd5b50565b600081519050611b3781611b11565b92915050565b600060208284031215611b5357611b5261102a565b5b6000611b6184828501611b28565b91505092915050565b600080fd5b600080fd5b600080fd5b60008083356001602003843603038112611b9657611b95611b6a565b5b80840192508235915067ffffffffffffffff821115611bb857611bb7611b6f565b5b602083019250600182023603831315611bd457611bd3611b74565b5b509250929050565b6000611be88385611994565b9350611bf58385846110fb565b82840190509392505050565b6000611c0e828486611bdc565b91508190509392505050565b7f7265766572740000000000000000000000000000000000000000000000000000600082015250565b6000611c506006836119e7565b9150611c5b82611c1a565b600682019050919050565b6000611c7182611c43565b9150819050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6000611cb582611312565b9150611cc083611312565b925082611cd057611ccf611c7b565b5b82820490509291505056fea2646970667358221220c65e0c220554f807998419d73152c256b7cda9add6caf7336e420daca02bf0cf64736f6c634300081a0033",
}

// TestDAppV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use TestDAppV2MetaData.ABI instead.
var TestDAppV2ABI = TestDAppV2MetaData.ABI

// TestDAppV2Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TestDAppV2MetaData.Bin instead.
var TestDAppV2Bin = TestDAppV2MetaData.Bin

// DeployTestDAppV2 deploys a new Ethereum contract, binding an instance of TestDAppV2 to it.
func DeployTestDAppV2(auth *bind.TransactOpts, backend bind.ContractBackend, isZetaChain_ bool, gateway_ common.Address) (common.Address, *types.Transaction, *TestDAppV2, error) {
	parsed, err := TestDAppV2MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TestDAppV2Bin), backend, isZetaChain_, gateway_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TestDAppV2{TestDAppV2Caller: TestDAppV2Caller{contract: contract}, TestDAppV2Transactor: TestDAppV2Transactor{contract: contract}, TestDAppV2Filterer: TestDAppV2Filterer{contract: contract}}, nil
}

// TestDAppV2 is an auto generated Go binding around an Ethereum contract.
type TestDAppV2 struct {
	TestDAppV2Caller     // Read-only binding to the contract
	TestDAppV2Transactor // Write-only binding to the contract
	TestDAppV2Filterer   // Log filterer for contract events
}

// TestDAppV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type TestDAppV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestDAppV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type TestDAppV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestDAppV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestDAppV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestDAppV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestDAppV2Session struct {
	Contract     *TestDAppV2       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestDAppV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestDAppV2CallerSession struct {
	Contract *TestDAppV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// TestDAppV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestDAppV2TransactorSession struct {
	Contract     *TestDAppV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// TestDAppV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type TestDAppV2Raw struct {
	Contract *TestDAppV2 // Generic contract binding to access the raw methods on
}

// TestDAppV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestDAppV2CallerRaw struct {
	Contract *TestDAppV2Caller // Generic read-only contract binding to access the raw methods on
}

// TestDAppV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestDAppV2TransactorRaw struct {
	Contract *TestDAppV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewTestDAppV2 creates a new instance of TestDAppV2, bound to a specific deployed contract.
func NewTestDAppV2(address common.Address, backend bind.ContractBackend) (*TestDAppV2, error) {
	contract, err := bindTestDAppV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestDAppV2{TestDAppV2Caller: TestDAppV2Caller{contract: contract}, TestDAppV2Transactor: TestDAppV2Transactor{contract: contract}, TestDAppV2Filterer: TestDAppV2Filterer{contract: contract}}, nil
}

// NewTestDAppV2Caller creates a new read-only instance of TestDAppV2, bound to a specific deployed contract.
func NewTestDAppV2Caller(address common.Address, caller bind.ContractCaller) (*TestDAppV2Caller, error) {
	contract, err := bindTestDAppV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestDAppV2Caller{contract: contract}, nil
}

// NewTestDAppV2Transactor creates a new write-only instance of TestDAppV2, bound to a specific deployed contract.
func NewTestDAppV2Transactor(address common.Address, transactor bind.ContractTransactor) (*TestDAppV2Transactor, error) {
	contract, err := bindTestDAppV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestDAppV2Transactor{contract: contract}, nil
}

// NewTestDAppV2Filterer creates a new log filterer instance of TestDAppV2, bound to a specific deployed contract.
func NewTestDAppV2Filterer(address common.Address, filterer bind.ContractFilterer) (*TestDAppV2Filterer, error) {
	contract, err := bindTestDAppV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestDAppV2Filterer{contract: contract}, nil
}

// bindTestDAppV2 binds a generic wrapper to an already deployed contract.
func bindTestDAppV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TestDAppV2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestDAppV2 *TestDAppV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestDAppV2.Contract.TestDAppV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestDAppV2 *TestDAppV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestDAppV2.Contract.TestDAppV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestDAppV2 *TestDAppV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestDAppV2.Contract.TestDAppV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestDAppV2 *TestDAppV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestDAppV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestDAppV2 *TestDAppV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestDAppV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestDAppV2 *TestDAppV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestDAppV2.Contract.contract.Transact(opts, method, params...)
}

// NOMESSAGECALL is a free data retrieval call binding the contract method 0xc85f8434.
//
// Solidity: function NO_MESSAGE_CALL() view returns(string)
func (_TestDAppV2 *TestDAppV2Caller) NOMESSAGECALL(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TestDAppV2.contract.Call(opts, &out, "NO_MESSAGE_CALL")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// NOMESSAGECALL is a free data retrieval call binding the contract method 0xc85f8434.
//
// Solidity: function NO_MESSAGE_CALL() view returns(string)
func (_TestDAppV2 *TestDAppV2Session) NOMESSAGECALL() (string, error) {
	return _TestDAppV2.Contract.NOMESSAGECALL(&_TestDAppV2.CallOpts)
}

// NOMESSAGECALL is a free data retrieval call binding the contract method 0xc85f8434.
//
// Solidity: function NO_MESSAGE_CALL() view returns(string)
func (_TestDAppV2 *TestDAppV2CallerSession) NOMESSAGECALL() (string, error) {
	return _TestDAppV2.Contract.NOMESSAGECALL(&_TestDAppV2.CallOpts)
}

// AmountWithMessage is a free data retrieval call binding the contract method 0x4297a263.
//
// Solidity: function amountWithMessage(bytes32 ) view returns(uint256)
func (_TestDAppV2 *TestDAppV2Caller) AmountWithMessage(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _TestDAppV2.contract.Call(opts, &out, "amountWithMessage", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AmountWithMessage is a free data retrieval call binding the contract method 0x4297a263.
//
// Solidity: function amountWithMessage(bytes32 ) view returns(uint256)
func (_TestDAppV2 *TestDAppV2Session) AmountWithMessage(arg0 [32]byte) (*big.Int, error) {
	return _TestDAppV2.Contract.AmountWithMessage(&_TestDAppV2.CallOpts, arg0)
}

// AmountWithMessage is a free data retrieval call binding the contract method 0x4297a263.
//
// Solidity: function amountWithMessage(bytes32 ) view returns(uint256)
func (_TestDAppV2 *TestDAppV2CallerSession) AmountWithMessage(arg0 [32]byte) (*big.Int, error) {
	return _TestDAppV2.Contract.AmountWithMessage(&_TestDAppV2.CallOpts, arg0)
}

// CalledWithMessage is a free data retrieval call binding the contract method 0xe2842ed7.
//
// Solidity: function calledWithMessage(bytes32 ) view returns(bool)
func (_TestDAppV2 *TestDAppV2Caller) CalledWithMessage(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _TestDAppV2.contract.Call(opts, &out, "calledWithMessage", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CalledWithMessage is a free data retrieval call binding the contract method 0xe2842ed7.
//
// Solidity: function calledWithMessage(bytes32 ) view returns(bool)
func (_TestDAppV2 *TestDAppV2Session) CalledWithMessage(arg0 [32]byte) (bool, error) {
	return _TestDAppV2.Contract.CalledWithMessage(&_TestDAppV2.CallOpts, arg0)
}

// CalledWithMessage is a free data retrieval call binding the contract method 0xe2842ed7.
//
// Solidity: function calledWithMessage(bytes32 ) view returns(bool)
func (_TestDAppV2 *TestDAppV2CallerSession) CalledWithMessage(arg0 [32]byte) (bool, error) {
	return _TestDAppV2.Contract.CalledWithMessage(&_TestDAppV2.CallOpts, arg0)
}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_TestDAppV2 *TestDAppV2Caller) Gateway(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TestDAppV2.contract.Call(opts, &out, "gateway")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_TestDAppV2 *TestDAppV2Session) Gateway() (common.Address, error) {
	return _TestDAppV2.Contract.Gateway(&_TestDAppV2.CallOpts)
}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_TestDAppV2 *TestDAppV2CallerSession) Gateway() (common.Address, error) {
	return _TestDAppV2.Contract.Gateway(&_TestDAppV2.CallOpts)
}

// GetAmountWithMessage is a free data retrieval call binding the contract method 0x9291fe26.
//
// Solidity: function getAmountWithMessage(string message) view returns(uint256)
func (_TestDAppV2 *TestDAppV2Caller) GetAmountWithMessage(opts *bind.CallOpts, message string) (*big.Int, error) {
	var out []interface{}
	err := _TestDAppV2.contract.Call(opts, &out, "getAmountWithMessage", message)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmountWithMessage is a free data retrieval call binding the contract method 0x9291fe26.
//
// Solidity: function getAmountWithMessage(string message) view returns(uint256)
func (_TestDAppV2 *TestDAppV2Session) GetAmountWithMessage(message string) (*big.Int, error) {
	return _TestDAppV2.Contract.GetAmountWithMessage(&_TestDAppV2.CallOpts, message)
}

// GetAmountWithMessage is a free data retrieval call binding the contract method 0x9291fe26.
//
// Solidity: function getAmountWithMessage(string message) view returns(uint256)
func (_TestDAppV2 *TestDAppV2CallerSession) GetAmountWithMessage(message string) (*big.Int, error) {
	return _TestDAppV2.Contract.GetAmountWithMessage(&_TestDAppV2.CallOpts, message)
}

// GetCalledWithMessage is a free data retrieval call binding the contract method 0xf592cbfb.
//
// Solidity: function getCalledWithMessage(string message) view returns(bool)
func (_TestDAppV2 *TestDAppV2Caller) GetCalledWithMessage(opts *bind.CallOpts, message string) (bool, error) {
	var out []interface{}
	err := _TestDAppV2.contract.Call(opts, &out, "getCalledWithMessage", message)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetCalledWithMessage is a free data retrieval call binding the contract method 0xf592cbfb.
//
// Solidity: function getCalledWithMessage(string message) view returns(bool)
func (_TestDAppV2 *TestDAppV2Session) GetCalledWithMessage(message string) (bool, error) {
	return _TestDAppV2.Contract.GetCalledWithMessage(&_TestDAppV2.CallOpts, message)
}

// GetCalledWithMessage is a free data retrieval call binding the contract method 0xf592cbfb.
//
// Solidity: function getCalledWithMessage(string message) view returns(bool)
func (_TestDAppV2 *TestDAppV2CallerSession) GetCalledWithMessage(message string) (bool, error) {
	return _TestDAppV2.Contract.GetCalledWithMessage(&_TestDAppV2.CallOpts, message)
}

// GetNoMessageIndex is a free data retrieval call binding the contract method 0xad23b28b.
//
// Solidity: function getNoMessageIndex(address sender) pure returns(string)
func (_TestDAppV2 *TestDAppV2Caller) GetNoMessageIndex(opts *bind.CallOpts, sender common.Address) (string, error) {
	var out []interface{}
	err := _TestDAppV2.contract.Call(opts, &out, "getNoMessageIndex", sender)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetNoMessageIndex is a free data retrieval call binding the contract method 0xad23b28b.
//
// Solidity: function getNoMessageIndex(address sender) pure returns(string)
func (_TestDAppV2 *TestDAppV2Session) GetNoMessageIndex(sender common.Address) (string, error) {
	return _TestDAppV2.Contract.GetNoMessageIndex(&_TestDAppV2.CallOpts, sender)
}

// GetNoMessageIndex is a free data retrieval call binding the contract method 0xad23b28b.
//
// Solidity: function getNoMessageIndex(address sender) pure returns(string)
func (_TestDAppV2 *TestDAppV2CallerSession) GetNoMessageIndex(sender common.Address) (string, error) {
	return _TestDAppV2.Contract.GetNoMessageIndex(&_TestDAppV2.CallOpts, sender)
}

// IsZetaChain is a free data retrieval call binding the contract method 0xc91f3567.
//
// Solidity: function isZetaChain() view returns(bool)
func (_TestDAppV2 *TestDAppV2Caller) IsZetaChain(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _TestDAppV2.contract.Call(opts, &out, "isZetaChain")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsZetaChain is a free data retrieval call binding the contract method 0xc91f3567.
//
// Solidity: function isZetaChain() view returns(bool)
func (_TestDAppV2 *TestDAppV2Session) IsZetaChain() (bool, error) {
	return _TestDAppV2.Contract.IsZetaChain(&_TestDAppV2.CallOpts)
}

// IsZetaChain is a free data retrieval call binding the contract method 0xc91f3567.
//
// Solidity: function isZetaChain() view returns(bool)
func (_TestDAppV2 *TestDAppV2CallerSession) IsZetaChain() (bool, error) {
	return _TestDAppV2.Contract.IsZetaChain(&_TestDAppV2.CallOpts)
}

// SenderWithMessage is a free data retrieval call binding the contract method 0xf936ae85.
//
// Solidity: function senderWithMessage(bytes ) view returns(address)
func (_TestDAppV2 *TestDAppV2Caller) SenderWithMessage(opts *bind.CallOpts, arg0 []byte) (common.Address, error) {
	var out []interface{}
	err := _TestDAppV2.contract.Call(opts, &out, "senderWithMessage", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SenderWithMessage is a free data retrieval call binding the contract method 0xf936ae85.
//
// Solidity: function senderWithMessage(bytes ) view returns(address)
func (_TestDAppV2 *TestDAppV2Session) SenderWithMessage(arg0 []byte) (common.Address, error) {
	return _TestDAppV2.Contract.SenderWithMessage(&_TestDAppV2.CallOpts, arg0)
}

// SenderWithMessage is a free data retrieval call binding the contract method 0xf936ae85.
//
// Solidity: function senderWithMessage(bytes ) view returns(address)
func (_TestDAppV2 *TestDAppV2CallerSession) SenderWithMessage(arg0 []byte) (common.Address, error) {
	return _TestDAppV2.Contract.SenderWithMessage(&_TestDAppV2.CallOpts, arg0)
}

// Erc20Call is a paid mutator transaction binding the contract method 0xc7a339a9.
//
// Solidity: function erc20Call(address erc20, uint256 amount, string message) returns()
func (_TestDAppV2 *TestDAppV2Transactor) Erc20Call(opts *bind.TransactOpts, erc20 common.Address, amount *big.Int, message string) (*types.Transaction, error) {
	return _TestDAppV2.contract.Transact(opts, "erc20Call", erc20, amount, message)
}

// Erc20Call is a paid mutator transaction binding the contract method 0xc7a339a9.
//
// Solidity: function erc20Call(address erc20, uint256 amount, string message) returns()
func (_TestDAppV2 *TestDAppV2Session) Erc20Call(erc20 common.Address, amount *big.Int, message string) (*types.Transaction, error) {
	return _TestDAppV2.Contract.Erc20Call(&_TestDAppV2.TransactOpts, erc20, amount, message)
}

// Erc20Call is a paid mutator transaction binding the contract method 0xc7a339a9.
//
// Solidity: function erc20Call(address erc20, uint256 amount, string message) returns()
func (_TestDAppV2 *TestDAppV2TransactorSession) Erc20Call(erc20 common.Address, amount *big.Int, message string) (*types.Transaction, error) {
	return _TestDAppV2.Contract.Erc20Call(&_TestDAppV2.TransactOpts, erc20, amount, message)
}

// GasCall is a paid mutator transaction binding the contract method 0xa799911f.
//
// Solidity: function gasCall(string message) payable returns()
func (_TestDAppV2 *TestDAppV2Transactor) GasCall(opts *bind.TransactOpts, message string) (*types.Transaction, error) {
	return _TestDAppV2.contract.Transact(opts, "gasCall", message)
}

// GasCall is a paid mutator transaction binding the contract method 0xa799911f.
//
// Solidity: function gasCall(string message) payable returns()
func (_TestDAppV2 *TestDAppV2Session) GasCall(message string) (*types.Transaction, error) {
	return _TestDAppV2.Contract.GasCall(&_TestDAppV2.TransactOpts, message)
}

// GasCall is a paid mutator transaction binding the contract method 0xa799911f.
//
// Solidity: function gasCall(string message) payable returns()
func (_TestDAppV2 *TestDAppV2TransactorSession) GasCall(message string) (*types.Transaction, error) {
	return _TestDAppV2.Contract.GasCall(&_TestDAppV2.TransactOpts, message)
}

// GatewayCall is a paid mutator transaction binding the contract method 0x41a3cd4a.
//
// Solidity: function gatewayCall(address dst, bytes payload) returns()
func (_TestDAppV2 *TestDAppV2Transactor) GatewayCall(opts *bind.TransactOpts, dst common.Address, payload []byte) (*types.Transaction, error) {
	return _TestDAppV2.contract.Transact(opts, "gatewayCall", dst, payload)
}

// GatewayCall is a paid mutator transaction binding the contract method 0x41a3cd4a.
//
// Solidity: function gatewayCall(address dst, bytes payload) returns()
func (_TestDAppV2 *TestDAppV2Session) GatewayCall(dst common.Address, payload []byte) (*types.Transaction, error) {
	return _TestDAppV2.Contract.GatewayCall(&_TestDAppV2.TransactOpts, dst, payload)
}

// GatewayCall is a paid mutator transaction binding the contract method 0x41a3cd4a.
//
// Solidity: function gatewayCall(address dst, bytes payload) returns()
func (_TestDAppV2 *TestDAppV2TransactorSession) GatewayCall(dst common.Address, payload []byte) (*types.Transaction, error) {
	return _TestDAppV2.Contract.GatewayCall(&_TestDAppV2.TransactOpts, dst, payload)
}

// GatewayDeposit is a paid mutator transaction binding the contract method 0x9ca016ed.
//
// Solidity: function gatewayDeposit(address dst) payable returns()
func (_TestDAppV2 *TestDAppV2Transactor) GatewayDeposit(opts *bind.TransactOpts, dst common.Address) (*types.Transaction, error) {
	return _TestDAppV2.contract.Transact(opts, "gatewayDeposit", dst)
}

// GatewayDeposit is a paid mutator transaction binding the contract method 0x9ca016ed.
//
// Solidity: function gatewayDeposit(address dst) payable returns()
func (_TestDAppV2 *TestDAppV2Session) GatewayDeposit(dst common.Address) (*types.Transaction, error) {
	return _TestDAppV2.Contract.GatewayDeposit(&_TestDAppV2.TransactOpts, dst)
}

// GatewayDeposit is a paid mutator transaction binding the contract method 0x9ca016ed.
//
// Solidity: function gatewayDeposit(address dst) payable returns()
func (_TestDAppV2 *TestDAppV2TransactorSession) GatewayDeposit(dst common.Address) (*types.Transaction, error) {
	return _TestDAppV2.Contract.GatewayDeposit(&_TestDAppV2.TransactOpts, dst)
}

// GatewayDepositAndCall is a paid mutator transaction binding the contract method 0xdeb3b1e4.
//
// Solidity: function gatewayDepositAndCall(address dst, bytes payload) payable returns()
func (_TestDAppV2 *TestDAppV2Transactor) GatewayDepositAndCall(opts *bind.TransactOpts, dst common.Address, payload []byte) (*types.Transaction, error) {
	return _TestDAppV2.contract.Transact(opts, "gatewayDepositAndCall", dst, payload)
}

// GatewayDepositAndCall is a paid mutator transaction binding the contract method 0xdeb3b1e4.
//
// Solidity: function gatewayDepositAndCall(address dst, bytes payload) payable returns()
func (_TestDAppV2 *TestDAppV2Session) GatewayDepositAndCall(dst common.Address, payload []byte) (*types.Transaction, error) {
	return _TestDAppV2.Contract.GatewayDepositAndCall(&_TestDAppV2.TransactOpts, dst, payload)
}

// GatewayDepositAndCall is a paid mutator transaction binding the contract method 0xdeb3b1e4.
//
// Solidity: function gatewayDepositAndCall(address dst, bytes payload) payable returns()
func (_TestDAppV2 *TestDAppV2TransactorSession) GatewayDepositAndCall(dst common.Address, payload []byte) (*types.Transaction, error) {
	return _TestDAppV2.Contract.GatewayDepositAndCall(&_TestDAppV2.TransactOpts, dst, payload)
}

// OnCall is a paid mutator transaction binding the contract method 0x5bcfd616.
//
// Solidity: function onCall((bytes,address,uint256) _context, address _zrc20, uint256 amount, bytes message) returns()
func (_TestDAppV2 *TestDAppV2Transactor) OnCall(opts *bind.TransactOpts, _context TestDAppV2zContext, _zrc20 common.Address, amount *big.Int, message []byte) (*types.Transaction, error) {
	return _TestDAppV2.contract.Transact(opts, "onCall", _context, _zrc20, amount, message)
}

// OnCall is a paid mutator transaction binding the contract method 0x5bcfd616.
//
// Solidity: function onCall((bytes,address,uint256) _context, address _zrc20, uint256 amount, bytes message) returns()
func (_TestDAppV2 *TestDAppV2Session) OnCall(_context TestDAppV2zContext, _zrc20 common.Address, amount *big.Int, message []byte) (*types.Transaction, error) {
	return _TestDAppV2.Contract.OnCall(&_TestDAppV2.TransactOpts, _context, _zrc20, amount, message)
}

// OnCall is a paid mutator transaction binding the contract method 0x5bcfd616.
//
// Solidity: function onCall((bytes,address,uint256) _context, address _zrc20, uint256 amount, bytes message) returns()
func (_TestDAppV2 *TestDAppV2TransactorSession) OnCall(_context TestDAppV2zContext, _zrc20 common.Address, amount *big.Int, message []byte) (*types.Transaction, error) {
	return _TestDAppV2.Contract.OnCall(&_TestDAppV2.TransactOpts, _context, _zrc20, amount, message)
}

// OnCall0 is a paid mutator transaction binding the contract method 0x676cc054.
//
// Solidity: function onCall((address) messageContext, bytes message) payable returns(bytes)
func (_TestDAppV2 *TestDAppV2Transactor) OnCall0(opts *bind.TransactOpts, messageContext TestDAppV2MessageContext, message []byte) (*types.Transaction, error) {
	return _TestDAppV2.contract.Transact(opts, "onCall0", messageContext, message)
}

// OnCall0 is a paid mutator transaction binding the contract method 0x676cc054.
//
// Solidity: function onCall((address) messageContext, bytes message) payable returns(bytes)
func (_TestDAppV2 *TestDAppV2Session) OnCall0(messageContext TestDAppV2MessageContext, message []byte) (*types.Transaction, error) {
	return _TestDAppV2.Contract.OnCall0(&_TestDAppV2.TransactOpts, messageContext, message)
}

// OnCall0 is a paid mutator transaction binding the contract method 0x676cc054.
//
// Solidity: function onCall((address) messageContext, bytes message) payable returns(bytes)
func (_TestDAppV2 *TestDAppV2TransactorSession) OnCall0(messageContext TestDAppV2MessageContext, message []byte) (*types.Transaction, error) {
	return _TestDAppV2.Contract.OnCall0(&_TestDAppV2.TransactOpts, messageContext, message)
}

// OnRevert is a paid mutator transaction binding the contract method 0xc9028a36.
//
// Solidity: function onRevert((address,address,uint256,bytes) revertContext) returns()
func (_TestDAppV2 *TestDAppV2Transactor) OnRevert(opts *bind.TransactOpts, revertContext TestDAppV2RevertContext) (*types.Transaction, error) {
	return _TestDAppV2.contract.Transact(opts, "onRevert", revertContext)
}

// OnRevert is a paid mutator transaction binding the contract method 0xc9028a36.
//
// Solidity: function onRevert((address,address,uint256,bytes) revertContext) returns()
func (_TestDAppV2 *TestDAppV2Session) OnRevert(revertContext TestDAppV2RevertContext) (*types.Transaction, error) {
	return _TestDAppV2.Contract.OnRevert(&_TestDAppV2.TransactOpts, revertContext)
}

// OnRevert is a paid mutator transaction binding the contract method 0xc9028a36.
//
// Solidity: function onRevert((address,address,uint256,bytes) revertContext) returns()
func (_TestDAppV2 *TestDAppV2TransactorSession) OnRevert(revertContext TestDAppV2RevertContext) (*types.Transaction, error) {
	return _TestDAppV2.Contract.OnRevert(&_TestDAppV2.TransactOpts, revertContext)
}

// SimpleCall is a paid mutator transaction binding the contract method 0x36e980a0.
//
// Solidity: function simpleCall(string message) returns()
func (_TestDAppV2 *TestDAppV2Transactor) SimpleCall(opts *bind.TransactOpts, message string) (*types.Transaction, error) {
	return _TestDAppV2.contract.Transact(opts, "simpleCall", message)
}

// SimpleCall is a paid mutator transaction binding the contract method 0x36e980a0.
//
// Solidity: function simpleCall(string message) returns()
func (_TestDAppV2 *TestDAppV2Session) SimpleCall(message string) (*types.Transaction, error) {
	return _TestDAppV2.Contract.SimpleCall(&_TestDAppV2.TransactOpts, message)
}

// SimpleCall is a paid mutator transaction binding the contract method 0x36e980a0.
//
// Solidity: function simpleCall(string message) returns()
func (_TestDAppV2 *TestDAppV2TransactorSession) SimpleCall(message string) (*types.Transaction, error) {
	return _TestDAppV2.Contract.SimpleCall(&_TestDAppV2.TransactOpts, message)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TestDAppV2 *TestDAppV2Transactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestDAppV2.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TestDAppV2 *TestDAppV2Session) Receive() (*types.Transaction, error) {
	return _TestDAppV2.Contract.Receive(&_TestDAppV2.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TestDAppV2 *TestDAppV2TransactorSession) Receive() (*types.Transaction, error) {
	return _TestDAppV2.Contract.Receive(&_TestDAppV2.TransactOpts)
}
