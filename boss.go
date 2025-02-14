package main


type BOSS_Header struct {
    magic1  [4]byte
    magic2  [4]byte
    Bigend  uint16
    serial  uint64
    unknown uint8
    ___     uint8 // Padding?
    hash    [32]byte
    rsa        [256]byte
    IV      [12]byte
}
type Content_Header struct {
    unknown    []byte
    payload 
    hash    [32]byte
    rsa        [256]byte
}
type Payload_Content_Header {
    ProgramID    
    unknown        
    Content        
    payload_size    
    NsDataId
    version
    hash    [32]byte
    rsa        [256]byte
}