package guac

import "testing"

func TestEncodeString(t *testing.T) {
	encoded := EncodeString("The \"&\" and \"<\" and \">\" should be encoded")
	expected := "The \"&amp;\" and \"&lt;\" and \"&gt;\" should be encoded"
	if encoded != expected {
		t.Fatal("Encoded string does not match", expected, encoded)
	}
}

func TestEncodeStringMultipleInstances(t *testing.T) {
	encoded := EncodeString("& < > & < > & < >")
	expected := "&amp; &lt; &gt; &amp; &lt; &gt; &amp; &lt; &gt;"
	if encoded != expected {
		t.Fatal("Encoded string does not match", expected, encoded)
	}
}

func TestDecodeString(t *testing.T) {
	decoded := DecodeString("The \"&amp;\" and \"&lt;\" and \"&gt;\" should be encoded")
	expected := "The \"&\" and \"<\" and \">\" should be encoded"
	if decoded != expected {
		t.Fatal("Decoded string does not match", expected, decoded)
	}
}

func TestDecodeStringMultipleInstances(t *testing.T) {
	decoded := DecodeString("&amp; &lt; &gt; &amp; &lt; &gt; &amp; &lt; &gt;")
	expected := "& < > & < > & < >"
	if decoded != expected {
		t.Fatal("Decoded string does not match", expected, decoded)
	}
}
