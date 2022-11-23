package provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccComplexResource(t *testing.T) {
	t.Cleanup(CleanupTestingDirectories(t))
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: ProviderFactories(""),
		Steps: []resource.TestStep{
			{
				Config: LoadFile(t, "testdata/complex/create/main.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("tfcoremock_complex_resource.test", "string", "hello"),
					resource.TestCheckResourceAttr("tfcoremock_complex_resource.test", "list.#", "1"),
					resource.TestCheckResourceAttr("tfcoremock_complex_resource.test", "list.0.string", "one"),
					resource.TestCheckResourceAttr("tfcoremock_complex_resource.test", "object.bool", "true"),
					resource.TestCheckResourceAttr("tfcoremock_complex_resource.test", "set.#", "2"),
					resource.TestCheckResourceAttr("tfcoremock_complex_resource.test", "set.0.string", "one"),
					resource.TestCheckResourceAttr("tfcoremock_complex_resource.test", "set.1.string", "zero")),
			},
			{
				Config: LoadFile(t, "testdata/complex/update/main.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("tfcoremock_complex_resource.test", "string", "hello"),
					resource.TestCheckResourceAttr("tfcoremock_complex_resource.test", "list.#", "2"),
					resource.TestCheckResourceAttr("tfcoremock_complex_resource.test", "list.0.string", "zero"),
					resource.TestCheckResourceAttr("tfcoremock_complex_resource.test", "list.1.string", "one"),
					resource.TestCheckResourceAttr("tfcoremock_complex_resource.test", "object.string", "world"),
					resource.TestCheckResourceAttr("tfcoremock_complex_resource.test", "set.#", "1"),
					resource.TestCheckResourceAttr("tfcoremock_complex_resource.test", "set.0.string", "zero")),
			},
			{
				Config: LoadFile(t, "testdata/complex/delete/main.tf"),
			},
		},
	})
}

func TestAccComplexResourceWithBlocks(t *testing.T) {
	t.Cleanup(CleanupTestingDirectories(t))
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: ProviderFactories(""),
		Steps: []resource.TestStep{
			{
				Config: LoadFile(t, "testdata/complex_block/create/main.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("tfcoremock_complex_resource.test", "list_block.#", "2"),
					resource.TestCheckResourceAttr("tfcoremock_complex_resource.test", "list_block.0.integer", "0"),
					resource.TestCheckResourceAttr("tfcoremock_complex_resource.test", "list_block.1.integer", "1"),
					resource.TestCheckResourceAttr("tfcoremock_complex_resource.test", "set_block.#", "1"),
					resource.TestCheckResourceAttr("tfcoremock_complex_resource.test", "set_block.0.integer", "0")),
			},
			{
				Config: LoadFile(t, "testdata/complex_block/update/main.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("tfcoremock_complex_resource.test", "list_block.#", "1"),
					resource.TestCheckResourceAttr("tfcoremock_complex_resource.test", "list_block.0.integer", "0"),
					resource.TestCheckResourceAttr("tfcoremock_complex_resource.test", "set_block.#", "2"),
					resource.TestCheckResourceAttr("tfcoremock_complex_resource.test", "set_block.0.integer", "0"),
					resource.TestCheckResourceAttr("tfcoremock_complex_resource.test", "set_block.1.integer", "1")),
			},
			{
				Config: LoadFile(t, "testdata/complex/delete/main.tf"),
			},
		},
	})
}

func TestAccDynamicResource(t *testing.T) {
	t.Cleanup(CleanupTestingDirectories(t))
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: ProviderFactories(LoadFile(t, "testdata/dynamic/dynamic_resources.json")),
		Steps: []resource.TestStep{
			{
				Config: LoadFile(t, "testdata/dynamic/create/main.tf"),
				Check:  resource.TestCheckResourceAttr("tfcoremock_dynamic_resource.test", "integer", "0"),
			},
			{
				Config: LoadFile(t, "testdata/dynamic/update/main.tf"),
				Check:  resource.TestCheckResourceAttr("tfcoremock_dynamic_resource.test", "integer", "1"),
			},
			{
				Config: LoadFile(t, "testdata/dynamic/delete/main.tf"),
			},
		},
	})
}

func TestAccDynamicResourceNested(t *testing.T) {
	t.Cleanup(CleanupTestingDirectories(t))
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: ProviderFactories(LoadFile(t, "testdata/dynamic_nested/dynamic_resources.json")),
		Steps: []resource.TestStep{
			{
				Config: LoadFile(t, "testdata/dynamic_nested/create/main.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("tfcoremock_dynamic_resource.test", "string", "hello"),
					resource.TestCheckResourceAttr("tfcoremock_dynamic_resource.test", "list.#", "1"),
					resource.TestCheckResourceAttr("tfcoremock_dynamic_resource.test", "list.0.string", "one"),
					resource.TestCheckResourceAttr("tfcoremock_dynamic_resource.test", "object.bool", "true")),
			},
			{
				Config: LoadFile(t, "testdata/dynamic/delete/main.tf"),
			},
		},
	})
}

func TestAccDynamicResourceWithBlocks(t *testing.T) {
	t.Cleanup(CleanupTestingDirectories(t))
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: ProviderFactories(LoadFile(t, "testdata/dynamic_block/dynamic_resources.json")),
		Steps: []resource.TestStep{
			{
				Config: LoadFile(t, "testdata/dynamic_block/create/main.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("tfcoremock_dynamic_resource.test", "integer", "0"),
					resource.TestCheckResourceAttr("tfcoremock_dynamic_resource.test", "nested_list.#", "1"),
					resource.TestCheckResourceAttr("tfcoremock_dynamic_resource.test", "nested_list.0.integer", "0")),
			},
			{
				Config: LoadFile(t, "testdata/dynamic_block/update/main.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("tfcoremock_dynamic_resource.test", "integer", "0"),
					resource.TestCheckResourceAttr("tfcoremock_dynamic_resource.test", "nested_list.#", "2"),
					resource.TestCheckResourceAttr("tfcoremock_dynamic_resource.test", "nested_list.0.integer", "0"),
					resource.TestCheckResourceAttr("tfcoremock_dynamic_resource.test", "nested_list.1.integer", "1")),
			},
			{
				Config: LoadFile(t, "testdata/dynamic/delete/main.tf"),
			},
		},
	})
}

func TestAccDynamicResourceWithSingleBlock(t *testing.T) {
	t.Cleanup(CleanupTestingDirectories(t))
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: ProviderFactories(LoadFile(t, "testdata/dynamic_single_block/dynamic_resources.json")),
		Steps: []resource.TestStep{
			{
				Config: LoadFile(t, "testdata/dynamic_single_block/create/main.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("tfcoremock_dynamic_resource.test", "integer", "0"),
					resource.TestCheckResourceAttr("tfcoremock_dynamic_resource.test", "nested_object.integer", "0")),
			},
			{
				Config: LoadFile(t, "testdata/dynamic_single_block/update/main.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("tfcoremock_dynamic_resource.test", "integer", "0"),
					resource.TestCheckResourceAttr("tfcoremock_dynamic_resource.test", "nested_object.integer", "1")),
			},
			{
				Config: LoadFile(t, "testdata/dynamic/delete/main.tf"),
			},
		},
	})
}

func TestAccDynamicResourceWithComputed(t *testing.T) {
	t.Cleanup(CleanupTestingDirectories(t))
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: ProviderFactories(LoadFile(t, "testdata/dynamic_computed/dynamic_resources.json")),
		Steps: []resource.TestStep{
			{
				Config: LoadFile(t, "testdata/dynamic_computed/create/main.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("tfcoremock_dynamic_resource.test", "object_with_value.boolean", "true"),
					resource.TestCheckResourceAttr("tfcoremock_dynamic_resource.test", "object_with_value.string", "hello"),
					resource.TestCheckResourceAttrSet("tfcoremock_dynamic_resource.test", "computed_object.id"),
					resource.TestCheckResourceAttr("tfcoremock_dynamic_resource.test", "computed_list.#", "0"),
					resource.TestCheckResourceAttr("tfcoremock_dynamic_resource.test", "set.#", "3"),
					resource.TestCheckResourceAttrSet("tfcoremock_dynamic_resource.test", "set.0.id"),
					resource.TestCheckResourceAttrSet("tfcoremock_dynamic_resource.test", "set.1.id"),
					resource.TestCheckResourceAttrSet("tfcoremock_dynamic_resource.test", "set.2.id")),
			},
			{
				Config: LoadFile(t, "testdata/dynamic/delete/main.tf"),
			},
		},
	})
}

func TestAccDynamicResourceWithComputedBlocks(t *testing.T) {
	t.Cleanup(CleanupTestingDirectories(t))
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: ProviderFactories(LoadFile(t, "testdata/dynamic_computed_block/dynamic_resources.json")),
		Steps: []resource.TestStep{
			{
				Config: LoadFile(t, "testdata/dynamic_computed_block/create/main.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("tfcoremock_dynamic_resource.test", "other.#", "2"),
					resource.TestCheckResourceAttr("tfcoremock_dynamic_resource.test", "other.0.id", "my-id"),
					resource.TestCheckResourceAttr("tfcoremock_dynamic_resource.test", "other.0.nested.#", "2"),
					resource.TestCheckResourceAttrSet("tfcoremock_dynamic_resource.test", "other.0.nested.0.id"),
					resource.TestCheckResourceAttrSet("tfcoremock_dynamic_resource.test", "other.0.nested.1.id"),
					resource.TestCheckResourceAttrSet("tfcoremock_dynamic_resource.test", "other.1.id"),
					resource.TestCheckResourceAttr("tfcoremock_dynamic_resource.test", "object.#", "0")),
			},
			{
				Config: LoadFile(t, "testdata/dynamic/delete/main.tf"),
			},
		},
	})
}

func TestAccDynamicResourceWithComputedSetBlocks(t *testing.T) {
	// TODO(liamcervante): Investigate and enable this test case.
	// This is the same as above with the nested block being a set instead of
	// a list. I don't think there's any difference in the way my provider
	// handles this so I think it's a bug in the framework.
	// https://github.com/hashicorp/terraform-plugin-framework/issues/483
	t.Skip()

	t.Cleanup(CleanupTestingDirectories(t))
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: ProviderFactories(LoadFile(t, "testdata/dynamic_computed_block_set/dynamic_resources.json")),
		Steps: []resource.TestStep{
			{
				Config: LoadFile(t, "testdata/dynamic_computed_block_set/create/main.tf"),
			},
			{
				Config: LoadFile(t, "testdata/dynamic/delete/main.tf"),
			},
		},
	})
}

func TestAccDynamicResourceWithId(t *testing.T) {
	t.Cleanup(CleanupTestingDirectories(t))
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: ProviderFactories(LoadFile(t, "testdata/dynamic/dynamic_resources.json")),
		Steps: []resource.TestStep{
			{
				Config: LoadFile(t, "testdata/dynamic_with_id/create/main.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("tfcoremock_dynamic_resource.test", "integer", "0"),
					resource.TestCheckResourceAttr("tfcoremock_dynamic_resource.test", "id", "my_id")),
			},
			{
				Config: LoadFile(t, "testdata/dynamic_with_id/update/main.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("tfcoremock_dynamic_resource.test", "integer", "1"),
					resource.TestCheckResourceAttr("tfcoremock_dynamic_resource.test", "id", "my_id")),
			},
			{
				Config: LoadFile(t, "testdata/dynamic/delete/main.tf"),
			},
		},
	})
}

func TestAccDynamicResourceWithDataSource(t *testing.T) {
	t.Cleanup(CleanupTestingDirectories(t))
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: ProviderFactories(LoadFile(t, "testdata/dynamic_datasource/dynamic_resources.json")),
		Steps: []resource.TestStep{
			{
				Config: LoadFile(t, "testdata/dynamic_datasource/create/main.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.tfcoremock_simple_resource.test", "integer", "0"),
					resource.TestCheckResourceAttr("tfcoremock_dynamic_resource.test", "id", "my_dynamic_resource"),
					resource.TestCheckResourceAttr("tfcoremock_dynamic_resource.test", "my_value", "0")),
			},
			{
				Config: LoadFile(t, "testdata/dynamic_datasource/delete/main.tf"),
			},
		},
	})
}

func TestAccSimpleDataSource(t *testing.T) {
	t.Cleanup(CleanupTestingDirectories(t))
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: ProviderFactories(""),
		Steps: []resource.TestStep{
			{
				Config: LoadFile(t, "testdata/simple_datasource/get/main.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.tfcoremock_simple_resource.data", "integer", "0"),
					resource.TestCheckResourceAttr("tfcoremock_simple_resource.test", "integer", "0")),
			},
		},
	})
}

func TestAccSimpleResource(t *testing.T) {
	t.Cleanup(CleanupTestingDirectories(t))
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: ProviderFactories(""),
		Steps: []resource.TestStep{
			{
				Config: LoadFile(t, "testdata/simple/create/main.tf"),
				Check:  resource.TestCheckResourceAttr("tfcoremock_simple_resource.test", "integer", "0"),
			},
			{
				Config: LoadFile(t, "testdata/simple/update/main.tf"),
				Check:  resource.TestCheckResourceAttr("tfcoremock_simple_resource.test", "integer", "1"),
			},
			{
				Config: LoadFile(t, "testdata/simple/delete/main.tf"),
			},
		},
	})
}

func TestAccSimpleResourceWithDrift(t *testing.T) {
	t.Cleanup(CleanupTestingDirectories(t))
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: ProviderFactories(""),
		Steps: []resource.TestStep{
			{
				Config: LoadFile(t, "testdata/simple/create/main.tf"),
				Check: func(state *terraform.State) error {
					id := state.RootModule().Resources["tfcoremock_simple_resource.test"].Primary.Attributes["id"]
					return os.Remove(fmt.Sprintf("terraform.resource/%s.json", id))
				},
				ExpectNonEmptyPlan: true,
			},
			{
				Config: LoadFile(t, "testdata/simple/update/main.tf"),
				Check:  resource.TestCheckResourceAttr("tfcoremock_simple_resource.test", "integer", "1"),
			},
			{
				Config: LoadFile(t, "testdata/simple/delete/main.tf"),
			},
		},
	})
}

func TestAccSimpleResourceWithId(t *testing.T) {
	t.Cleanup(CleanupTestingDirectories(t))
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: ProviderFactories(""),
		Steps: []resource.TestStep{
			{
				Config: LoadFile(t, "testdata/simple_with_id/create/main.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("tfcoremock_simple_resource.test", "string", "hello"),
					resource.TestCheckResourceAttr("tfcoremock_simple_resource.test", "id", "my_id")),
			},
			{
				Config: LoadFile(t, "testdata/simple_with_id/update/main.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("tfcoremock_simple_resource.test", "string", "world"),
					resource.TestCheckResourceAttr("tfcoremock_simple_resource.test", "id", "my_id")),
			},
			{
				Config: LoadFile(t, "testdata/simple/delete/main.tf"),
			},
		},
	})
}

func TestAccRemoteResource(t *testing.T) {
	t.Cleanup(CleanupTestingDirectories(t))
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: ProviderFactories(LoadFile(t, "testdata/remote/dynamic_resources.json")),
		Steps: []resource.TestStep{
			{
				Config: LoadFile(t, "testdata/remote/create/main.tf"),
				Check:  resource.TestCheckResourceAttrSet("null_resource.test", "id"),
			},
			{
				Config: LoadFile(t, "testdata/remote/delete/main.tf"),
			},
		},
	})
}
