<?php
// DO NOT EDIT! Generated by Protobuf-PHP protoc plugin 1.0
// Source: query.proto
//   Date: 2016-01-22 01:34:42

namespace Vitess\Proto\Query {

  class Row extends \DrSlump\Protobuf\Message {

    /**  @var int[]  */
    public $lengths = array();
    
    /**  @var string */
    public $values = null;
    

    /** @var \Closure[] */
    protected static $__extensions = array();

    public static function descriptor()
    {
      $descriptor = new \DrSlump\Protobuf\Descriptor(__CLASS__, 'query.Row');

      // REPEATED SINT64 lengths = 1
      $f = new \DrSlump\Protobuf\Field();
      $f->number    = 1;
      $f->name      = "lengths";
      $f->type      = \DrSlump\Protobuf::TYPE_SINT64;
      $f->rule      = \DrSlump\Protobuf::RULE_REPEATED;
      $descriptor->addField($f);

      // OPTIONAL BYTES values = 2
      $f = new \DrSlump\Protobuf\Field();
      $f->number    = 2;
      $f->name      = "values";
      $f->type      = \DrSlump\Protobuf::TYPE_BYTES;
      $f->rule      = \DrSlump\Protobuf::RULE_OPTIONAL;
      $descriptor->addField($f);

      foreach (self::$__extensions as $cb) {
        $descriptor->addField($cb(), true);
      }

      return $descriptor;
    }

    /**
     * Check if <lengths> has a value
     *
     * @return boolean
     */
    public function hasLengths(){
      return $this->_has(1);
    }
    
    /**
     * Clear <lengths> value
     *
     * @return \Vitess\Proto\Query\Row
     */
    public function clearLengths(){
      return $this->_clear(1);
    }
    
    /**
     * Get <lengths> value
     *
     * @param int $idx
     * @return int
     */
    public function getLengths($idx = NULL){
      return $this->_get(1, $idx);
    }
    
    /**
     * Set <lengths> value
     *
     * @param int $value
     * @return \Vitess\Proto\Query\Row
     */
    public function setLengths( $value, $idx = NULL){
      return $this->_set(1, $value, $idx);
    }
    
    /**
     * Get all elements of <lengths>
     *
     * @return int[]
     */
    public function getLengthsList(){
     return $this->_get(1);
    }
    
    /**
     * Add a new element to <lengths>
     *
     * @param int $value
     * @return \Vitess\Proto\Query\Row
     */
    public function addLengths( $value){
     return $this->_add(1, $value);
    }
    
    /**
     * Check if <values> has a value
     *
     * @return boolean
     */
    public function hasValues(){
      return $this->_has(2);
    }
    
    /**
     * Clear <values> value
     *
     * @return \Vitess\Proto\Query\Row
     */
    public function clearValues(){
      return $this->_clear(2);
    }
    
    /**
     * Get <values> value
     *
     * @return string
     */
    public function getValues(){
      return $this->_get(2);
    }
    
    /**
     * Set <values> value
     *
     * @param string $value
     * @return \Vitess\Proto\Query\Row
     */
    public function setValues( $value){
      return $this->_set(2, $value);
    }
  }
}

