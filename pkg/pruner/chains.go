package pruner

type Chain struct {
	name string
	keys []string
}

var chains = []Chain{
	{
		name: "osmosis",
		keys: []string{
			"icahost",        //icahosttypes.StoreKey,
			"gamm",           // gammtypes.StoreKey,
			"lockup",         //lockuptypes.StoreKey,
			"incentives",     // incentivestypes.StoreKey,
			"epochs",         // epochstypes.StoreKey,
			"poolincentives", //poolincentivestypes.StoreKey,
			"txfees",         // txfeestypes.StoreKey,
			"superfluid",     // superfluidtypes.StoreKey,
			"bech32ibc",      // bech32ibctypes.StoreKey,
			"wasm",           // wasm.StoreKey,
			"tokenfactory",   //tokenfactorytypes.StoreKey,
		},
	},
	{
		name: "cosmoshub",
		keys: []string{
			"liquidity",
			"icahost", // icahosttypes.StoreKey
		},
	},
	{
		name: "terra",
		keys: []string{
			"oracle",   // oracletypes.StoreKey,
			"market",   // markettypes.StoreKey,
			"treasury", //treasurytypes.StoreKey,
			"wasm",     // wasmtypes.StoreKey,
		},
	},
	{
		name: "kava",
		keys: []string{
			"feemarket", //feemarkettypes.StoreKey,
			"kavadist",  //kavadisttypes.StoreKey,
			"auction",   //auctiontypes.StoreKey,
			"issuance",  //issuancetypes.StoreKey,
			"bep3",      //bep3types.StoreKey,
			"cdp",       //cdptypes.StoreKey,
			"hard",      //hardtypes.StoreKey,
			"committee", //committeetypes.StoreKey,
			"incentive", //incentivetypes.StoreKey,
			"evmutil",   //evmutiltypes.StoreKey,
			"savings",   //savingstypes.StoreKey,
			"bridge",    //bridgetypes.StoreKey,
		},
	},
	{
		name: "evmos",
		keys: []string{
			"evm",        // evmtypes.StoreKey,
			"feemarket",  // feemarkettypes.StoreKey,
			"inflation",  // inflationtypes.StoreKey,
			"erc20",      // erc20types.StoreKey,
			"incentives", // incentivestypes.StoreKey,
			"epochs",     // epochstypes.StoreKey,
			"claims",     // claimstypes.StoreKey,
			"vesting",    // vestingtypes.StoreKey,
		},
	},
	{
		name: "akash",
		keys: []string{
			"escrow",     // escrow.StoreKey,
			"deployment", // deployment.StoreKey,
			"market",     // market.StoreKey,
			"provider",   // provider.StoreKey,
			"audit",      // audit.StoreKey,
			"cert",       // cert.StoreKey,
			"inflation",  // inflation.StoreKey,
		},
	},
	{
		name: "juno",
		keys: []string{
			"icahost", // icahosttypes.StoreKey,
			"wasm",    // wasm.StoreKey,
		},
	},
	{
		name: "stargaze",
		keys: []string{
			"claim", // claimmoduletypes.StoreKey,
			"alloc", // allocmoduletypes.StoreKey,
			"wasm",  // wasm.StoreKey,
		},
	},
	{
		name: "omniflix",
		keys: []string{
			"alloc",       // alloctypes.StoreKey,
			"onft",        // onfttypes.StoreKey,
			"marketplace", // marketplacetypes.StoreKey,
		},
	},
	{
		name: "irisnet",
		keys: []string{
			"guardian", // guardiantypes.StoreKey,
			"token",    // tokentypes.StoreKey,
			"nft",      // nfttypes.StoreKey,
			"htlc",     // htlctypes.StoreKey,
			"record",   // recordtypes.StoreKey,
			"coinswap", // coinswaptypes.StoreKey,
			"service",  // servicetypes.StoreKey,
			"oracle",   // oracletypes.StoreKey,
			"random",   // randomtypes.StoreKey,
			"farm",     // farmtypes.StoreKey,
			"tibc",     // tibchost.StoreKey,
			"NFT",      // tibcnfttypes.StoreKey,
			"MT",       // tibcmttypes.StoreKey,
			"mt",       // mttypes.StoreKey,
		},
	},
	{
		name: "axelar",
		keys: []string{
			"vote",       // voteTypes.StoreKey,
			"evm",        // evmTypes.StoreKey,
			"snapshot",   // snapTypes.StoreKey,
			"multisig",   // multisigTypes.StoreKey,
			"tss",        // tssTypes.StoreKey,
			"nexus",      // nexusTypes.StoreKey,
			"axelarnet",  // axelarnetTypes.StoreKey,
			"reward",     // rewardTypes.StoreKey,
			"permission", // permissionTypes.StoreKey,
			"wasm",       // wasm.StoreKey,
		},
	},
}

func getChains() map[string][]string {
	chainsMap := make(map[string][]string)
	for _, c := range chains {
		chainsMap[c.name] = c.keys
	}

	return chainsMap
}
