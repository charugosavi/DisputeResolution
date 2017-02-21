/*
Licensed Materials - Property of IBM

6949-63F

(C) Copyright IBM Corp. 2016, 2016
All Rights Reserved

US Government Users Restricted Rights - Use, duplication or
disclosure restricted by GSA ADP Schedule Contract with IBM Corp.
*/

/* Do not modify this source code directly.
 * This is automatically generated by model_crud.go.php */

package main

import (
	"encoding/json"
	"fmt"
	"errors"
)

var _ = fmt.Printf
var _ = errors.New("Temp")

<?php
require_once(dirname(__FILE__)."/"."chaincode.php");

//FIXME treating comment
error_reporting(E_STRICT);

function terminate_missing_variables($errno, $errstr, $errfile, $errline)
{                               
  if (($errno == E_NOTICE) and (strstr($errstr, "Undefined variable")))
   die ("$errstr in $errfile line $errline");

  return false; // Let the PHP error handler handle all the rest  
}

$old_error_handler = set_error_handler("terminate_missing_variables"); 

$schema = loadSchema("dispute_management_schema.go");
$models = array_keys($schema);
?>

//List structures
<?php
foreach ($schema as $model => $definition) {
	if ($definition["listable"] && $model != "Reference") { ?>
type <?php echo $model; ?>s struct {
	Data []<?php echo $model; ?>
	
}	
<?php
	}
}
?>

type Dump struct {
<?php
foreach ($schema as $model => $definition) {
	if ($definition["listable"]) { ?>
	<?php echo $model; ?> *<?php echo $model; ?>s
<?php
	}
}
?>
}

func (this *HDLS) dump() (*Dump, error) {
	var d Dump
	var err error
<?php
foreach ($schema as $model => $definition) {
	if ($definition["listable"]) { ?>
	d.<?php echo $model; ?>, err = this.list<?php echo $model; ?>s()
	if err != nil {
		return nil, err
	}
<?php
	}
}
?>

	return &d, err
}

func (this *HDLS) imprt(dump *Dump) (error) {
	var err error
<?php
foreach ($schema as $model => $definition) {
	if ($definition["listable"]) { ?>
	if dump.<?php echo $model; ?> != nil {
		for _, x := range dump.<?php echo $model; ?>.Data {
			err = this.put<?php echo $model; ?>(&x)
		}
	}
<?php
	}
}
?>
	return err
}

func (this *HDLS) imprtJson(jsonStr string) (error) {
	var d Dump
	err := json.Unmarshal([]byte(jsonStr), &d)
	if err != nil {
		return err
	}

	return this.imprt(&d)
}

func (this *HDLS) createSchema() {
	models := []string{
<?php
foreach ($schema as $model => $definition) {
?>
		"<?php echo $model; ?>", 
<?php
}
?>
	}
	for _, model := range models {
		this.createKeyEntTable(model)
	}
}
<?php
$i = 1;
foreach ($schema as $model => $definition) {
?>
//------------------------
// <?php echo ($i++);?>. <?php echo strtoupper($model);?> 
//------------------------

<?php foreach ($definition["index"] as $field) { ?>
func (this *HDLS) refId<?php echo $model; ?><?php echo $field["name"]; ?>(v <?php echo $field["type"]; ?>) string {
	return fmt.Sprintf("<?php echo $model; ?>.<?php echo $field["name"]; ?>=%v", v)
}
<?php } ?>

func (this *HDLS) put<?php echo $model; ?>(x *<?php echo $model; ?>) error {
	this.logger.Infof("Call: put<?php echo $model; ?> with id: " + x.Id)
<?php  if ( $model != "InvokingStatus" ) { //FIXME magic varible ?>
	if x.Id == "" {
		x.Id, _ = this.id<?php echo $model; ?>(x)
		this.logger.Infof("Id is set to: " + x.Id)
	}
<?php  } ?>

	dst := x	// copy

<?php  foreach ($definition["relations"] as $relation) { ?>
	//Save dst.<?php echo $relation["element"]; ?> as a separate entity
	if dst.<?php echo $relation["element"]; ?> != nil {
		dst.<?php echo $relation["element"]; ?>.Id = dst.Id + "_<?php echo $relation["element"]; ?>"
		err1 := this.put<?php echo $relation["model"]; ?>(dst.<?php echo $relation["element"]; ?>)
		if err1 != nil {
			return err1
		}
	}
<?php }?>

	//Remove all the referenced entities since they are already stored.
<?php  foreach ($definition["relations"] as $relation) { ?>
	dst.<?php echo $relation["element"]; ?> = nil
<?php }?>	
	err := this.putA("<?php echo $model; ?>", dst.Id, dst)
	if err != nil {
		return err
	}
<?php if ($definition["index"]) { ?>
	var ref *Reference
	var refId string
<?php } ?>
<?php foreach ($definition["index"] as $field) { ?>
	refId = this.refId<?php echo $model; ?><?php echo $field["name"]; ?>(x.<?php echo $field["name"]; ?>)
	ref, _ = this.getReference(refId)
	if ref == nil {
		ref = &Reference{
			Id : refId,
			Ids: []string{dst.Id},
		}
		err = this.putReference(ref)
	} else {
		ref.Ids = append(ref.Ids, dst.Id)
		err = this.overwriteReference(ref)
	}
	if err != nil {
		return err
	}
	
<?php } ?>

	return nil
}

func (this *HDLS) get<?php echo $model; ?>(id string) (*<?php echo $model; ?>, error) {
	this.logger.Infof("Call: get<?php echo $model; ?> with id: " + id)

	var x <?php echo $model; ?> 
	err := this.getA("<?php echo $model; ?>", id, &x)
	if err != nil {
		this.logger.Infof("Error occured %v\n", err)
		return nil, err
	} else if x.Id == "" {
		return nil, nil
	}

<?php  foreach ($definition["relations"] as $relation) { ?>
	x.<?php echo $relation["element"]; ?>, err = this.get<?php echo $relation["model"]; ?>(x.Id + "_<?php echo $relation["element"]; ?>")
	if err != nil {
		return nil, err
	}
<?php }?>

	return &x, nil
}

<?php if ($definition["listable"]) { ?>
func (this *HDLS) list<?php echo $model; ?>s() (*<?php echo $model; ?>s, error) {
	this.logger.Infof("Call: list<?php echo $model; ?>")

	rows, err := this.listAllRows("<?php echo $model; ?>")
	if err != nil {
		return nil, err
	}

	var xs <?php echo $model; ?>s
	for _, row := range rows {
		var x <?php echo $model; ?> 
		if this.val(row, &x) == nil {
<?php  foreach ($definition["relations"] as $relation) { ?>
			x.<?php echo $relation["element"]; ?>, err = this.get<?php echo $relation["model"]; ?>(x.Id + "_<?php echo $relation["element"]; ?>")
			if err != nil {
				continue
			}
<?php }?>
<?php /* TODO: access control */?>
			xs.Data = append(xs.Data, x)
		}
	}
	return &xs, nil
}
<?php } ?>

<?php foreach ($definition["index"] as $field) { ?>
func (this *HDLS) list<?php echo $model; ?>sBy<?php echo $field["name"]; ?>(v <?php echo $field["type"]; ?>) (*<?php echo $model; ?>s, error) {

	var xs <?php echo $model; ?>s
	refId := this.refId<?php echo $model; ?><?php echo $field["name"]; ?>(v)
	reference, _ := this.getReference(refId)
	if reference != nil {
		for _, id := range reference.Ids {
			x, _ := this.get<?php echo $model; ?>(id)
			if x != nil {
				xs.Data = append(xs.Data, *x)
			}
		}
	}

	return &xs, nil
}
<?php } ?>

func (this *HDLS) add<?php echo $model; ?>(jsonStr string) error {

	var x <?php echo $model; ?> 
	err := json.Unmarshal([]byte(jsonStr), &x)
	if err != nil {
		return err
	}

	return this.put<?php echo $model; ?>(&x)
}

func (this *HDLS) id<?php echo $model; ?>(x *<?php echo $model; ?>) (string, error) {
<?php if ( $definition["PK"] ) { 
	$keys = array();
	foreach ( $definition["PK"] as $pk ) { 
		$keys[] = "x.$pk";
	}
	?>
	return <?php echo join(" + \":\" + ", $keys);?>, nil
<?php } else { ?>
	return this.db.GetTxID(), nil
<?php } ?>
}

func (this *HDLS) delete<?php echo $model; ?>(x *<?php echo $model; ?>) error {
	this.logger.Infof("Call: delete<?php echo $model; ?> with id: " + x.Id)
	var err error
<?php if ($definition["index"]) { ?>
	var ref *Reference
	var refId string	

	curr, err := this.get<?php echo $model; ?>(x.Id)
	if err != nil {
		return err
	} else if curr == nil {
		return nil
	}
<?php }?>
<?php foreach ($definition["index"] as $field) { ?>
	refId = this.refId<?php echo $model; ?><?php echo $field["name"]; ?>(curr.<?php echo $field["name"]; ?>)
	ref, _ = this.getReference(refId)
	if ref != nil {
		ref.Ids = remove(ref.Ids, x.Id)
		this.deleteReference(ref)
		
		if len(ref.Ids) > 0 {
			this.overwriteReference(ref)
		} else {
			this.deleteReference(ref)
		}
	}
<?php }?>
<?php  foreach ($definition["relations"] as $relation) { ?>
	//Delete x.<?php echo $relation["element"]; ?>
	
	if(x.<?php echo $relation["element"]; ?> != nil) {
		this.logger.Infof("Deleting x.<?php echo $relation["element"]; ?> with id: " + x.<?php echo $relation["element"]; ?>.Id)
		err = this.delete<?php echo $relation["model"]; ?>(x.<?php echo $relation["element"]; ?>)
		if err != nil {
			return err
		}
	}
<?php }?>
	err = this.delete("<?php echo $model; ?>", x.Id)
	if err != nil {
		return err
	}

	return nil
}

func (this *HDLS) overwrite<?php echo $model; ?>(x *<?php echo $model; ?>) error {
	if err := this.delete<?php echo $model; ?>(x); err != nil {
		return err
	}
	
	return this.put<?php echo $model; ?>(x)
}
<?php
}
?>
