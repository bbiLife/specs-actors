// Code generated by github.com/whyrusleeping/cbor-gen. DO NOT EDIT.

package builtin

import (
	"fmt"
	"io"

	abi "github.com/filecoin-project/go-state-types/abi"
	cbg "github.com/whyrusleeping/cbor-gen"
	xerrors "golang.org/x/xerrors"
)

var _ = xerrors.Errorf

var lengthBufConfirmSectorProofsParams = []byte{132}

func (t *ConfirmSectorProofsParams) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write(lengthBufConfirmSectorProofsParams); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.Sectors ([]abi.SectorNumber) (slice)
	if len(t.Sectors) > cbg.MaxLength {
		return xerrors.Errorf("Slice value in field t.Sectors was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajArray, uint64(len(t.Sectors))); err != nil {
		return err
	}
	for _, v := range t.Sectors {
		if err := cbg.CborWriteHeader(w, cbg.MajUnsignedInt, uint64(v)); err != nil {
			return err
		}
	}

	// t.RewardSmoothed (smoothing.FilterEstimate) (struct)
	if err := t.RewardSmoothed.MarshalCBOR(w); err != nil {
		return err
	}

	// t.RewardBaselinePower (big.Int) (struct)
	if err := t.RewardBaselinePower.MarshalCBOR(w); err != nil {
		return err
	}

	// t.QualityAdjPowerSmoothed (smoothing.FilterEstimate) (struct)
	if err := t.QualityAdjPowerSmoothed.MarshalCBOR(w); err != nil {
		return err
	}
	return nil
}

func (t *ConfirmSectorProofsParams) UnmarshalCBOR(r io.Reader) error {
	*t = ConfirmSectorProofsParams{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 4 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Sectors ([]abi.SectorNumber) (slice)

	maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("t.Sectors: array too large (%d)", extra)
	}

	if maj != cbg.MajArray {
		return fmt.Errorf("expected cbor array")
	}

	if extra > 0 {
		t.Sectors = make([]abi.SectorNumber, extra)
	}

	for i := 0; i < int(extra); i++ {

		maj, val, err := cbg.CborReadHeaderBuf(br, scratch)
		if err != nil {
			return xerrors.Errorf("failed to read uint64 for t.Sectors slice: %w", err)
		}

		if maj != cbg.MajUnsignedInt {
			return xerrors.Errorf("value read for array t.Sectors was not a uint, instead got %d", maj)
		}

		t.Sectors[i] = abi.SectorNumber(val)
	}

	// t.RewardSmoothed (smoothing.FilterEstimate) (struct)

	{

		if err := t.RewardSmoothed.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.RewardSmoothed: %w", err)
		}

	}
	// t.RewardBaselinePower (big.Int) (struct)

	{

		if err := t.RewardBaselinePower.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.RewardBaselinePower: %w", err)
		}

	}
	// t.QualityAdjPowerSmoothed (smoothing.FilterEstimate) (struct)

	{

		if err := t.QualityAdjPowerSmoothed.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.QualityAdjPowerSmoothed: %w", err)
		}

	}
	return nil
}

var lengthBufDeferredCronEventParams = []byte{131}

func (t *DeferredCronEventParams) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write(lengthBufDeferredCronEventParams); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.EventPayload ([]uint8) (slice)
	if len(t.EventPayload) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.EventPayload was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajByteString, uint64(len(t.EventPayload))); err != nil {
		return err
	}

	if _, err := w.Write(t.EventPayload[:]); err != nil {
		return err
	}

	// t.RewardSmoothed (smoothing.FilterEstimate) (struct)
	if err := t.RewardSmoothed.MarshalCBOR(w); err != nil {
		return err
	}

	// t.QualityAdjPowerSmoothed (smoothing.FilterEstimate) (struct)
	if err := t.QualityAdjPowerSmoothed.MarshalCBOR(w); err != nil {
		return err
	}
	return nil
}

func (t *DeferredCronEventParams) UnmarshalCBOR(r io.Reader) error {
	*t = DeferredCronEventParams{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 3 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.EventPayload ([]uint8) (slice)

	maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}

	if extra > cbg.ByteArrayMaxLen {
		return fmt.Errorf("t.EventPayload: byte array too large (%d)", extra)
	}
	if maj != cbg.MajByteString {
		return fmt.Errorf("expected byte array")
	}

	if extra > 0 {
		t.EventPayload = make([]uint8, extra)
	}

	if _, err := io.ReadFull(br, t.EventPayload[:]); err != nil {
		return err
	}
	// t.RewardSmoothed (smoothing.FilterEstimate) (struct)

	{

		if err := t.RewardSmoothed.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.RewardSmoothed: %w", err)
		}

	}
	// t.QualityAdjPowerSmoothed (smoothing.FilterEstimate) (struct)

	{

		if err := t.QualityAdjPowerSmoothed.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.QualityAdjPowerSmoothed: %w", err)
		}

	}
	return nil
}
