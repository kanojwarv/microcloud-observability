# Object Model

Objects:
- Cluster
- Host
- VM
- Project
- Storage Pool
- Ceph OSD
- Ceph Pool
- Network

Relationships:
Cluster -> Host -> VM
Cluster -> Ceph -> OSD -> Pool
