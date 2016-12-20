<?php
/*
Licensed Materials - Property of IBM

6949-63F

(C) Copyright IBM Corp. 2016, 2016
All Rights Reserved

US Government Users Restricted Rights - Use, duplication or
disclosure restricted by GSA ADP Schedule Contract with IBM Corp.
*/

function loadSchema($source) {
	$fp = fopen($source, "r");
	$prev_model = null;

	$schema = array(
		"Reference" => array(
			"listable" => true, "PK"=>array(), "relations"=>array(), "index"=>array()
		)
	);

    $ident = "[A-Za-z0-9_]+";

	while (!feof($fp)) {
		$line = trim(fgets($fp));

		// Tables
		if ( $prev_model && 
			preg_match("/type\\s+({$prev_model}s)\\s+struct/", $line, $matches)) {

			$schema[$prev_model]["listable"] = true;
		}
		else if (preg_match("#type\\s+({$ident})\\s+struct.*@dominant#", $line, $matches)) {
		}
		else if (preg_match("/type\\s+({$ident})\\s+struct/", $line, $matches)) {
			list($_, $model) = $matches;
			$schema[$model] = array(	
					"listable" => false, "PK"=>array(), "relations"=>array(), "index"=>array());
			$prev_model = $model;
		}

		// Fields
		if ( $prev_model ) {
			if ( preg_match("#({$ident})\\s+({$ident}).*@PK#", $line, $matches)) {	
			$schema[$prev_model]["PK"][] = $matches[1];
			}
			if ( preg_match("#({$ident})\\s+({$ident}).*@index#", $line, $matches)) {	
				$schema[$prev_model]["index"][] = array("name"=>$matches[1], "type"=>$matches[2]);
			}
			if ( preg_match("#({$ident})\\s+\*({$ident})#", $line, $matches)) {	
				if ( $prev_model == "Trade" ) continue;// FIXME magic variable

				$schema[$prev_model]["relations"][] = $matches[1];
			}

			if ( preg_match("#\\s+(Name)\\s+({$ident})#", $line, $matches)) {	
				$schema[$prev_model]["title"] = $matches[1];
			}
			else {
				$schema[$prev_model]["title"] = "Id";
			}
		}

	}
	return $schema;
}

function queryChaincode($address, $user, $function, $mode="query", $arg0="10001") {
    $id = CC_ID;
    $req_id = $mode == "query" ? 5 : 3;
    $arg0 = json_encode($arg0);
    $data = <<<EOS
{
  "jsonrpc": "2.0",
  "method": "$mode",
  "params": {
      "type": 1,
      "chaincodeID":{
          "name":"$id"
      },
      "ctorMsg": {
         "function":"$function",
         "args":[$arg0]
      },
      "secureContext": "$user"
  },
  "id": $req_id
}
EOS;
    echo $data."\n";
	$options = array(
	    'http' => array(
	        'header'  => "Content-type: application/x-www-form-urlencoded\r\n",
	        'method'  => 'POST',
	        'content' => $data
	    )
	);
	$context  = stream_context_create($options);
	$result = file_get_contents("http://$address/chaincode", false, $context);
	if ($result === FALSE) { /* Handle error */ }
	
	$result = json_decode($result, true);

    if ( $mode == "query") {
        return json_decode($result["result"]["message"], true);
    }
    else {
        usleep(1000*500);
        return $result["result"]["message"];
    }
}

?>
