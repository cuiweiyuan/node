package legacy

import (
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	"github.com/zeta-chain/node/e2e/runner"
	"github.com/zeta-chain/node/e2e/utils"
	"github.com/zeta-chain/node/testutil/sample"
)

func TestZetaDepositRestricted(r *runner.E2ERunner, args []string) {
	require.Len(r, args, 1)

	// parse the deposit amount
	amount := utils.ParseBigInt(r, args[0])

	// Deposit amount to restricted address
	txHash := r.LegacyDepositZetaWithAmount(ethcommon.HexToAddress(sample.RestrictedEVMAddressTest), amount)

	// wait for 5 zeta blocks
	r.WaitForBlocks(5)

	// no cctx should be created
	utils.EnsureNoCctxMinedByInboundHash(r.Ctx, txHash.String(), r.CctxClient)
}