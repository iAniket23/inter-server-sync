package schemareader

func applyTableFilters(table Table) Table {
	switch table.Name {
	case "rhnchecksumtype":
		table.PKSequence = "rhn_checksum_id_seq"
	case "rhnpackagearch":
		table.PKSequence = "rhn_package_arch_id_seq"
	case "rhnchannelarch":
		table.PKSequence = "rhn_channel_arch_id_seq"
	case "rhnpackagename":
		// constraint: rhn_pn_id_pk
		table.PKSequence = "RHN_PKG_NAME_SEQ"
	case "rhnpackageevr":
		// constraint: rhn_pe_id_pk
		table.PKSequence = "rhn_pkg_evr_seq"
	case "rhnpackage":
		// We need to add a virtual unique constraint
		table.PKSequence = "RHN_PACKAGE_ID_SEQ"
		virtualIndexName := "virtual_main_unique_index"
		virtualIndexColumns := []string{"name_id", "evr_id", "package_arch_id", "checksum_id", "org_id"}
		table.UniqueIndexes[virtualIndexName] = UniqueIndex{Name: virtualIndexName, Columns: virtualIndexColumns}
		table.MainUniqueIndexName = virtualIndexName
	case "rhnpackagecapability":
		// pkid: rhn_pkg_capability_id_pk
		table.PKSequence = "RHN_PKG_CAPABILITY_ID_SEQ"
	case "suseproductchannel": //FIXME we should try to add a unique constraint to this table instead of this hack
		// We need to add a virtual unique constraint
		virtualIndexName := "virtual_main_unique_index"
		virtualIndexColumns := []string{"product_id", "channel_id"}
		table.UniqueIndexes[virtualIndexName] = UniqueIndex{Name: virtualIndexName, Columns: virtualIndexColumns}
		table.MainUniqueIndexName = virtualIndexName
	case "susemdkeyword": //FIXME we should try to add a unique constraint to this table instead of this hack
		// We need to add a virtual unique constraint
		virtualIndexName := "virtual_main_unique_index"
		virtualIndexColumns := []string{"label"}
		table.UniqueIndexes[virtualIndexName] = UniqueIndex{Name: virtualIndexName, Columns: virtualIndexColumns}
		table.MainUniqueIndexName = virtualIndexName
	}
	return table
}
