# rubrik-exporter
Rubrik metrics exporter for Prometheus

Commandline Options
=====================

        Usage of rubrik-exporter:
        -listen-address string
                The address to lisiten on for HTTP requests. (default ":9477")
        -log.level value
                Only log messages with the given severity or above. Valid levels: [debug, info, warn, error, fatal, panic].
        -rubrik.password string
                Zerto API User Password
        -rubrik.url string
                Rubrik URL to connect https://rubrik.local.host
        -rubrik.username string
                Zerto API User

Environment variables can be used in lieu of commandline flags by uppercasing the
entire flag (e.g. you can ommit -rubrik.password if RUBRIK.PASSWORD is set)

Easy to use
=============

Using Docker 

    docker run --detached --publish 9477:9477 claranet/rubrik-exporter \
                          -rubrik.url https://myrubrik.company.org \
                          -rubrik.username "prometheus@local" \
                          -rubrik.password 'SeCure'

Exported Metrics
==================

        # HELP rubrik_archive_location_status Archive Loction Status - 1: Active, 0: Inactive
        # TYPE rubrik_archive_location_status gauge
        rubrik_archive_location_status{bucket="archive",name="NFS:archive",target="<ip-address>"} 1
        # HELP rubrik_archive_storage_archived_fileset ...
        # TYPE rubrik_archive_storage_archived_fileset gauge
        rubrik_archive_storage_archived_fileset{name="NFS:archive",target="<ip-address>",type="fileset"} 0
        rubrik_archive_storage_archived_fileset{name="NFS:archive",target="<ip-address>",type="linux"} 0
        rubrik_archive_storage_archived_fileset{name="NFS:archive",target="<ip-address>",type="share"} 0
        rubrik_archive_storage_archived_fileset{name="NFS:archive",target="<ip-address>",type="windows"} 0
        # HELP rubrik_archive_storage_archived_vm ...
        # TYPE rubrik_archive_storage_archived_vm gauge
        rubrik_archive_storage_archived_vm{name="NFS:archive",target="<ip-address>",type="hyperv"} 0
        rubrik_archive_storage_archived_vm{name="NFS:archive",target="<ip-address>",type="nutanix"} 0
        rubrik_archive_storage_archived_vm{name="NFS:archive",target="<ip-address>",type="vmware"} 0
        # HELP rubrik_archive_storage_data_archived ...
        # TYPE rubrik_archive_storage_data_archived gauge
        rubrik_archive_storage_data_archived{name="NFS:archive",target="<ip-address>"} 0
        # HELP rubrik_archive_storage_data_downloaded ...
        # TYPE rubrik_archive_storage_data_downloaded gauge
        rubrik_archive_storage_data_downloaded{name="NFS:archive",target="<ip-address>"} 0
        # HELP rubrik_count_nodes Count Rubrik Nodes in a Brick
        # TYPE rubrik_count_nodes gauge
        rubrik_count_nodes{brik="<brik-id>"} 4
        # HELP rubrik_count_streams Count Rubrik Backup Streams
        # TYPE rubrik_count_streams gauge
        rubrik_count_streams 0
        # HELP rubrik_node_io_read Node Read IO per second
        # TYPE rubrik_node_io_read gauge
        rubrik_node_io_read{node="<node-id>"} 281
        # HELP rubrik_node_io_write Node Write IO per second
        # TYPE rubrik_node_io_write gauge
        rubrik_node_io_write{node="<node-id>"} 1676
        # HELP rubrik_node_network_received Node Network Byte received
        # TYPE rubrik_node_network_received gauge
        rubrik_node_network_received{node="<node-id>"} 2.3782571e+07
        # HELP rubrik_node_network_transmitted Node Network Byte transmitted
        # TYPE rubrik_node_network_transmitted gauge
        rubrik_node_network_transmitted{node="<node-id>"} 2.0558578e+07
        # HELP rubrik_node_throughput_read Node Read Throughput per second
        # TYPE rubrik_node_throughput_read gauge
        rubrik_node_throughput_read{node="<node-id>"} 3.0397071e+07
        # HELP rubrik_node_throughput_write Node Write Throughput per second
        # TYPE rubrik_node_throughput_write gauge
        rubrik_node_throughput_write{node="<node-id>"} 2.5802489e+07
        # HELP rubrik_system_storage_available ...
        # TYPE rubrik_system_storage_available gauge
        rubrik_system_storage_available 7.0788234461184e+13
        # HELP rubrik_system_storage_live_mount ...
        # TYPE rubrik_system_storage_live_mount gauge
        rubrik_system_storage_live_mount 0
        # HELP rubrik_system_storage_miscellaneous ...
        # TYPE rubrik_system_storage_miscellaneous gauge
        rubrik_system_storage_miscellaneous 4.48470796529e+12
        # HELP rubrik_system_storage_snapshot ...
        # TYPE rubrik_system_storage_snapshot gauge
        rubrik_system_storage_snapshot 4.5390468081302e+13
        # HELP rubrik_system_storage_total ...
        # TYPE rubrik_system_storage_total gauge
        rubrik_system_storage_total 1.20663410507776e+14
        # HELP rubrik_system_storage_used ...
        # TYPE rubrik_system_storage_used gauge
        rubrik_system_storage_used 4.8594657722368e+13
        # HELP rubrik_vm_consumed_exclusive_bytes ...
        # TYPE rubrik_vm_consumed_exclusive_bytes gauge
        rubrik_vm_consumed_exclusive_bytes{vmname="<vm-name>""} 0
        # HELP rubrik_vm_consumed_index_storage_bytes ...
        # TYPE rubrik_vm_consumed_index_storage_bytes gauge
        rubrik_vm_consumed_index_storage_bytes{vmname="<vm-name>""} 0
        # HELP rubrik_vm_consumed_ingested_bytes ...
        # TYPE rubrik_vm_consumed_ingested_bytes gauge
        rubrik_vm_consumed_ingested_bytes{vmname="<vm-name>"} 0
        # HELP rubrik_vm_consumed_logical_bytes ...
        # TYPE rubrik_vm_consumed_logical_bytes gauge
        rubrik_vm_consumed_logical_bytes{vmname="<vm-name>""} 0
        # HELP rubrik_vm_consumed_shared_physical_bytes ...
        # TYPE rubrik_vm_consumed_shared_physical_bytes gauge
        rubrik_vm_consumed_shared_physical_bytes{vmname="<vm-name>"} 0
        # HELP rubrik_vm_protected ...
        # TYPE rubrik_vm_protected gauge
        rubrik_vm_protected{vmname="<vm-name>"} 0|1
