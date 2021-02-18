pragma solidity >=0.7.0 <0.8.0;
pragma abicoder v2;

contract Storage {
    
    mapping( string => string[] ) store;
    string[] tokens;

    function put(string memory _token, string[] memory _hashes) public {
        for(uint i=0;i<tokens.length;i++){
            if(keccak256(abi.encodePacked(tokens[i]))==keccak256(abi.encodePacked(_token))){
                for(uint j=0;j<_hashes.length;j++){
                    store[_token].push(_hashes[j]);
                }
                return;
            }
        }
        store[_token] = _hashes;
        tokens.push(_token);
        return;
    }
    
    function getTokens() public view returns (string[] memory) {
        return tokens;
    }
    
    function get(string memory _token) public view returns (string[] memory){
        return store[_token];
    }
    
}