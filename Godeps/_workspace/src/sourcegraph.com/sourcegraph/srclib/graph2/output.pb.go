// Code generated by protoc-gen-gogo.
// source: output.proto
// DO NOT EDIT!

package graph2

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// discarding unused import gogoproto "github.com/gogo/protobuf/gogoproto"
import ann "sourcegraph.com/sourcegraph/srclib/ann"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Output defines the output schema for the graph2 toolchain command.
// A srclib language toolchain's graph2 command should emit output in
// this format serialized as JSON.
type Output struct {
	// DefNodes is the list of definition nodes in the build unit's code graph.
	DefNodes []*Node `protobuf:"bytes,1,rep,name=def_nodes" json:"def_nodes,omitempty"`
	// RefNodes is the list of reference nodes in the build unit's code graph.
	RefNodes []*Node `protobuf:"bytes,2,rep,name=ref_nodes" json:"ref_nodes,omitempty"`
	// RefEdges is the list of edges (from reference nodes to definition
	// nodes) describing reference relationships in the source code
	// graph.
	RefEdges []*Edge `protobuf:"bytes,3,rep,name=ref_edges" json:"ref_edges,omitempty"`
	// DocNodes is the list of documentation nodes in the build unit's
	// code graph.
	DocNodes []*Node `protobuf:"bytes,4,rep,name=doc_nodes" json:"doc_nodes,omitempty"`
	// DocEdges is the list of edges tying documentation nodes to
	// definition nodes in the build unit's code graph.
	DocEdges []*Edge `protobuf:"bytes,5,rep,name=doc_edges" json:"doc_edges,omitempty"`
	// OtherNodes is the list of other nodes in the documentation's code
	// graph. For example, AST-level nodes.
	OtherNodes []*Node `protobuf:"bytes,6,rep,name=other_nodes" json:"other_nodes,omitempty"`
	// OtherEdges is the list of other edges in the build unit's code
	// graph. This can include type hierarchy, interface-implementation,
	// type-receiver, and AST-level relationships.
	OtherEdges []*Edge    `protobuf:"bytes,7,rep,name=other_edges" json:"other_edges,omitempty"`
	Anns       []*ann.Ann `protobuf:"bytes,8,rep,name=anns" json:"Anns,omitempty"`
}

func (m *Output) Reset()         { *m = Output{} }
func (m *Output) String() string { return proto.CompactTextString(m) }
func (*Output) ProtoMessage()    {}